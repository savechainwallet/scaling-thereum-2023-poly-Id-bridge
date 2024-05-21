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
      qrData:"{\"id\":\"c56a01d6-9614-4d00-878e-b730ef250dcf\",\"typ\":\"application/iden3comm-plain-json\",\"type\":\"https://iden3-communication.io/authorization/1.0/request\",\"thid\":\"c18645e2-aba3-4363-8396-7b83a2401b31\",\"from\":\"did:polyid:polygon:amoy:A74917835AF5A2E353fA40D8F73154ee3b4A6c02\",\"body\":{\"callbackUrl\":\"https://7xa5w6-ip-188-163-94-41.tunnelmole.net/callback?session_id=69e15ca5-dbbf-460c-ae7a-3c95430477ae\",\"reason\":\"KYC Verification\",\"message\":\"Verify Your KYC\",\"scope\":[{\"id\":1,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"date_of_birth\":null}}},{\"id\":2,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"address_country\":null}}},{\"id\":3,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"address_state\":null}}},{\"id\":4,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"address_unit\":null}}},{\"id\":5,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"address_street\":null}}},{\"id\":6,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"address_city\":null}}},{\"id\":7,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"address_zip\":null}}},{\"id\":7,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"email\":null}}},{\"id\":8,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"first_name\":null}}},{\"id\":9,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"last_name\":null}}},{\"id\":10,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"last_name\":null}}},{\"id\":11,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"middle_name\":null}}},{\"id\":12,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"nationality\":null}}},{\"id\":13,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"phone\":null}}},{\"id\":14,\"circuitId\":\"credentialAtomicQuerySigV2\",\"query\":{\"allowedIssuers\":[\"*\"],\"context\":\"https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/0idCredentials.jsonld\",\"type\":\"kyc\",\"credentialSubject\":{\"residence\":null}}}]}}",
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
          console.log(session.user)
          this.setSession(this.sessionId)
          this.$router.push({ path: 'userinfo' }).catch(err => {})

        }

      } catch (e) {
        console.error(e)

      }
      

    }

  },
  async created(){
   
   


  },
  beforeUnmount(){

  }
}
</script>

<style >

</style>