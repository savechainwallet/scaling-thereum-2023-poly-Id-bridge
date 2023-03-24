package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/models"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/utils"
	"gorm.io/gorm"
)

type Clients struct {
	db *gorm.DB
}

type ClientRegistrationRequest struct {
	Name        string `json:"name"`
	RedirectUrl string `json:"redirectUrl"`
}

func (c Clients) Register(ctx *gin.Context) {
	body := ClientRegistrationRequest{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Bad body", "error": err.Error()})
		return
	}
	did, key, err := utils.GenerateDID("did:polyid:polygon:mumbai")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Create DID error", "error": err.Error()})
		return
	}
	client := models.Client{
		Id:          did,
		Secret:      key,
		Name:        body.Name,
		RedirectUrl: body.RedirectUrl,
	}
	err = client.Save(c.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, client)

}
