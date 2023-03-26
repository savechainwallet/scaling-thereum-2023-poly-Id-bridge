<template lang="pug">
.fixed-top.d-flex.align-items-center.justify-content-center.centred
  b-card(border-variant="secondary" header="Users Pofile" title="Users profile is unconfirmed" style="max-width: 60rem;")
    b-card-text Please fill the data and confirm with KYC 
    b-form 
      b-container 
        b-row
          b-col(cols="4") 
            label(for="user-first-name") First Name
          b-col 
            b-form-input(v-model="firstName" id="user-first-name")
        b-row.mt-2
          b-col(cols="4") 
            label(for="user-last-name") Last Name
          b-col 
            b-form-input(v-model="lastName" id="user-last-name")
        b-row.mt-2
          b-col(cols="4") 
            label(for="user-email") Email
          b-col 
            b-form-input(v-model="email" id="user-email")
        b-row.mt-2
          b-col 
            b-button( variant="primary" @click="confirmationStart") Confirm
  b-modal(ref="qr-modal" hide-footer title="Scan for generate proof" size="lg")
    b-container 
      b-row(align-v="center")
        b-col
        b-col
          qr-code(:text="qrData"  color="#494b4b" bg-color="#4aba91" error-level="L" :size="500")
        b-col
    
      


</template>
  
  <script>
  import {GenerateProof,GetSession} from "../services/back"
  import {mapGetters, mapState, mapMutations} from "vuex"
  
  export default {
    name: 'HomeView',
    
    data(){
      return {
        firstName:"",
        lastName:"",
        email:"",
        qrData:"",
        timerId:""
      }
    },
    methods: {
      ...mapMutations(["setUser","setSession"]),
      async confirmationStart(){
        this.qrData = await GenerateProof({first_name:this.firstName, last_name:this.lastName, email:this.email}, this.sessionId)
        this.$refs['qr-modal'].show()
        let timerId = setInterval(this.confirmationEnd,1000)
        this.timerId = timerId

      },
      async confirmationEnd(){
      try {
        const session = await GetSession(this.sessionId)
        if (session.user.is_authorized){
          clearInterval(this.timerId)
          this.setUser(session.user)
          this.setSession(this.sessionId)
          window.location.href = this.getRedirectUrl
        }

      } catch (e) {
        console.error(e)

      }
      

    }

    },
    computed:{
      ...mapGetters(["userAuthorized","getRedirectUrl","getUserInfo"]),
      ...mapState(["sessionId"])

    },
    async created(){
      if (this.userAuthorized){
        window.location.href = this.getRedirectUrl
      } else {
        const user = this.getUserInfo
        this.firstName = user.first_name
        this.lastName = user.last_name
        this.email = user.email

      }
      
    
    }
  }
  </script>
  
  <style scoped >
  .centred{
  bottom: 0;
  overflow-y: auto;
}
  
  </style>