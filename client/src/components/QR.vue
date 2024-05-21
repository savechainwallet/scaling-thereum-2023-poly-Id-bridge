<template lang="pug">
.fixed-top.d-flex.align-items-center.justify-content-center.centred
  b-card(border-variant="secondary" header="Scan this on PolygonID App").text-center
    b-card-text
      qr-code( type="image/png" :color="{ dark: '#000000ff', light: '#ffffffff' }" value='{"id":"187db632-cad7-43a7-b9fb-3d5be94f398d","typ":"application/iden3comm-plain-json","type":"https://iden3-communication.io/authorization/1.0/request","thid":"40d8636a-2a4b-4fe1-9fc8-447d9d02e6f8","from":"did:polyid:polygon:amoy:A74917835AF5A2E353fA40D8F73154ee3b4A6c02","body":{"callbackUrl":"https://7xa5w6-ip-188-163-94-41.tunnelmole.net/callback?session_id=a5c47368-4e6b-44bd-a92d-28fdcaf075d3","reason":"KYC Verification","message":"Verify Your KYC","scope":[{"id":1,"circuitId":"credentialAtomicQuerySigV2","query":{"allowedIssuers":["*"],"context":"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld","type":"kyc","credentialSubject":{"date_of_birth":{}}}},{"id":2,"circuitId":"credentialAtomicQuerySigV2","query":{"allowedIssuers":["*"],"context":"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld","type":"kyc","credentialSubject":{"email":{}}}},{"id":3,"circuitId":"credentialAtomicQuerySigV2","query":{"allowedIssuers":["*"],"context":"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld","type":"kyc","credentialSubject":{"first_name":{}}}},{"id":4,"circuitId":"credentialAtomicQuerySigV2","query":{"allowedIssuers":["*"],"context":"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld","type":"kyc","credentialSubject":{"last_name":{}}}},{"id":5,"circuitId":"credentialAtomicQuerySigV2","query":{"allowedIssuers":["*"],"context":"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld","type":"kyc","credentialSubject":{"middle_name":{}}}},{"id":6,"circuitId":"credentialAtomicQuerySigV2","query":{"allowedIssuers":["*"],"context":"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld","type":"kyc","credentialSubject":{"phone":{}}}},{"id":7,"circuitId":"credentialAtomicQuerySigV2","query":{"allowedIssuers":["*"],"context":"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld","type":"kyc","credentialSubject":{"residence":{}}}}]}}')


      
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