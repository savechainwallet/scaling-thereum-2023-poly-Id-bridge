<template lang="pug">
.fixed-top.d-flex.align-items-center.justify-content-center.centred
  b-card(border-variant="secondary" header="Scan this on PolygonID App").text-center
    b-card-text
      qr-code(:text="qrData" :size="500" color="#494b4b" bg-color="#4aba91" error-level="L")


      
</template>

<script>
export default {
  name: "Login",
  props:["qr-data"],
  data(){
    return{
      
    }
  },

  
}
</script>

<style scoped>
.centred{
  bottom: 0;
  overflow-y: auto;
}
.login-card{
  max-width: 700px;
  min-width: 300px;
}
</style>