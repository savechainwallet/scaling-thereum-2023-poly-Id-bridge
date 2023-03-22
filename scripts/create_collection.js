import { Polybase } from '@polybase/client'
import { ethPersonalSign } from '@polybase/eth'

const db = new Polybase()

const PRIVATE_KEY = ''

db.signer((data) => {
  return {
    h: 'eth-personal-sign',
    sig: ethPersonalSign(PRIVATE_KEY, data)
  }
}
)

const createResponse = await db.applySchema(`
  collection ClientCollection {
    id: string;
    secret: string;
    name: string;
    redirectUrl: string;


    constructor (id: string, secret: string,name: string,redirectUrl: string) {
      this.id = id;
      this.secret = secret;
      this.name = name;
      this.redirectUrl = redirectUrl;
    }
  }
  collection UserCollection {
    id: string;
    firstName: string;
    lastName: string;
    country: string;
    isAccepted: boolean;
    clientId: string; 
    
    constructor(id: string, firstName: string, lastName: string, country: string, isAccepted: boolean, clientId: string) {
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.country = country;
        this.isAccepted = isAccepted;
        this.clientId = clientId;
    }
  }
  collection Session {
    id: string;
    userId: string;
    authRequest: bytes;
    isVerified: boolean;
    claims: bytes;

    constructor(id: string, userId: string, authRequest:bytes, isVerified: boolean, claims:bytes) {
        this.id = id;
        this.userId = userId;
        this.authRequest = authRequest;
        this.isVerified = isVerified;
        this.claims = claims;
        
    }
    del () {
        selfdestruct();
    }
  }


`,"poly-id-bridge") 

console.log(createResponse)