package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iden3/go-circuits"
	auth "github.com/iden3/go-iden3-auth"
	"github.com/iden3/go-iden3-auth/loaders"
	"github.com/iden3/go-iden3-auth/pubsignals"
	"github.com/iden3/go-iden3-auth/state"
	"github.com/iden3/iden3comm/protocol"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/models"
	"gorm.io/gorm"
)

type PolygonIDService struct {
	RedirectUrl        string
	RPCUrl             string
	ValidatorsContract string
	ResolverPrefix     string
	KeyDir             string
	db                 *gorm.DB
}

func (s *PolygonIDService) GetBasicQR(c *gin.Context) {

	reqId := uuid.NewString()
	session := models.Session{}
	session.RequestId = reqId
	if err := session.Save(s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error})
		return
	}
	if _, err := session.GetByRequestId(reqId, s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error})
		return
	}

	clientId := c.Query("clientId") // Eatch client has  did whitch use us audience
	client := models.Client{}
	_, err := client.GetById(clientId, s.db)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	request := auth.CreateAuthorizationRequest("bridgeauth", clientId, fmt.Sprintf("%s?session_id=%d", s.RedirectUrl, session.Id))
	request.ID = reqId
	reqJSON, _ := json.Marshal(request)
	session.AuthRequest = reqJSON
	session.RedirectUrl = client.RedirectUrl
	session.ClientId = client.Id
	if err := session.Save(s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"session_id": session.Id, "req": string(reqJSON)})

}

type CreateProofRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (s *PolygonIDService) GenerateProof(c *gin.Context) {

	body := CreateProofRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error})
		return
	}

	reqId := uuid.NewString()
	session := models.Session{}
	sessionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error})
		return
	}
	if _, err := session.GetById(uint(sessionId), s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error})
		return
	}

	var request protocol.AuthorizationRequestMessage = auth.CreateAuthorizationRequest("confirm-flow", session.ClientId, fmt.Sprintf("%s?session_id=%d", s.RedirectUrl, session.Id))
	request.ID = reqId

	var mtpProofRequest protocol.ZeroKnowledgeProofRequest
	mtpProofRequest.ID = 1
	mtpProofRequest.CircuitID = string(circuits.AtomicQuerySigV2CircuitID)
	mtpProofRequest.Query = map[string]interface{}{
		"allowedIssuers": []string{"*"},
		"credentialSubject": map[string]interface{}{
			"firstName": map[string]interface{}{
				"$eq": body.FirstName,
			},
		},
		"context": "https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/schemas/users-finance-kyc.jsonld",
		"type":    "UsersFinanceKYC",
	}
	var mtpProofRequest2 protocol.ZeroKnowledgeProofRequest
	mtpProofRequest2.ID = 2
	mtpProofRequest2.CircuitID = string(circuits.AtomicQuerySigV2CircuitID)
	mtpProofRequest2.Query = map[string]interface{}{
		"allowedIssuers": []string{"*"},
		"credentialSubject": map[string]interface{}{
			"lastName": map[string]interface{}{
				"$eq": body.LastName,
			},
		},
		"context": "https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/schemas/users-finance-kyc.jsonld",
		"type":    "UsersFinanceKYC",
	}

	var mtpProofRequest3 protocol.ZeroKnowledgeProofRequest
	mtpProofRequest3.ID = 3
	mtpProofRequest3.CircuitID = string(circuits.AtomicQuerySigV2CircuitID)
	mtpProofRequest3.Query = map[string]interface{}{
		"allowedIssuers": []string{"*"},
		"credentialSubject": map[string]interface{}{
			"emailAddress": map[string]interface{}{
				"$eq": body.Email,
			},
		},
		"context": "https://raw.githubusercontent.com/savechainwallet/scaling-thereum-2023-poly-Id-bridge/main/schemas/users-finance-kyc.jsonld",
		"type":    "UsersFinanceKYC",
	}
	request.Body.Scope = append(request.Body.Scope, mtpProofRequest)
	request.Body.Scope = append(request.Body.Scope, mtpProofRequest2)
	request.Body.Scope = append(request.Body.Scope, mtpProofRequest3)

	reqJSON, _ := json.Marshal(request)

	claimsJSON, _ := json.Marshal(body)
	session.AuthRequest = reqJSON
	session.Claims = claimsJSON
	if err := session.Save(s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, request)

}

func (s *PolygonIDService) ServeRedirect(c *gin.Context) {
	sesId, err := strconv.Atoi(c.Query("session_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	tokenBytes, _ := io.ReadAll(c.Request.Body)

	session := models.Session{}
	if _, err := session.GetById(uint(sesId), s.db); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var authRequest protocol.AuthorizationRequestMessage
	if err := json.Unmarshal(session.AuthRequest, &authRequest); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	verificationKeyloader := &loaders.FSKeyLoader{Dir: s.KeyDir}

	resolver := state.ETHResolver{
		RPCUrl:          s.RPCUrl,
		ContractAddress: common.HexToAddress(s.ValidatorsContract),
	}

	resolvers := map[string]pubsignals.StateResolver{
		s.ResolverPrefix: resolver,
	}

	verifier := auth.NewVerifier(verificationKeyloader, loaders.DefaultSchemaLoader{IpfsURL: "ipfs.io"}, resolvers)
	// authResponse, err := verifier.FullVerify(
	// 	c.Request.Context(),
	// 	string(tokenBytes),
	// 	authRequest)

	// vr

	t, err := verifier.VerifyJWZ(context.Background(), string(tokenBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	// parse jwz payload as json message
	var authResponse protocol.AuthorizationResponseMessage
	msg := t.GetPayload()
	err = json.Unmarshal(msg, &authResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	userId := authResponse.From

	session.UserId = userId
	if session.Connected {
		claims := CreateProofRequest{}
		if err := json.Unmarshal(session.Claims, &claims); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}
		user := models.User{}
		if _, err := user.GetById(userId, s.db); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}
		user.FirstName = claims.FirstName
		user.LastName = claims.LastName
		user.Email = claims.Email
		user.IsAccepted = true
		if err := user.Save(s.db); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

	} else {
		session.Connected = true
	}
	session.IsVerified = true
	if err := session.Save(s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"user": userId})

}
