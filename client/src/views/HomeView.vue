<template lang="pug">
QR(:qr-data="qrData")
</template>

<script>
import QR from '@/components/QR.vue'
import {GetQR, GetSession} from "../services/back"
import {mapMutations} from "vuex"

export default {
  name: 'HomeView',
  components: {
    QR
  },
  data(){
    return {
      qrData:"",
      sessionId:"",
      timerId:"",
    }
  },
  methods : {
    ...mapMutations(["setUser","setSession"]),
    async serveRedirect(){
      try {
        const session = await GetSession(this.sessionId)
        if (session.connected) {
          clearInterval(this.timerId)
          this.setUser(session.user)
          this.setSession(this.sessionId)
          this.$router.push({ path: 'userinfo' }).catch(err => {})

        }

      } catch (e) {
        console.error(e)

      }
      

    }

  },
  async created(){
   const queryString = window.location.search
   const urlParams = new URLSearchParams(queryString)
   const clientId = urlParams.get('client_id')
   const data = await GetQR(clientId)
   

   this.qrData = data.req
   this.sessionId = data.session_id

   let timerId = setInterval(this.serveRedirect,1000)
   this.timerId = timerId

  },
  beforeUnmount(){
    clearInterval(this.timerId)

  }
}
</script>

<style >

</style>