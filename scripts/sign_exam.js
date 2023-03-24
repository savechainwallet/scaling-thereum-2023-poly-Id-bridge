import { secp256k1 } from '@polybase/util'

// You can use any 32 byte private key from any library, 
// here is an example from @polybase/util
let privateKey = secp256k1.generatePrivateKey();
console.log(privateKey)
function createSig (timestamp, body) {
  const str_to_sign = `${timestamp}.${JSON.stringify(body)}`;
  return secp256k1.sign(privateKey, str_to_sign);
}

const sig = createSig("12345667784564",{"ff":"bb"})

console.log(sig)