package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/models"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/utils"
	"gorm.io/gorm"
)

type SessionController struct {
	db *gorm.DB
}

type SessionResponse struct {
	User         models.User `json:"user"`
	Code         string      `json:"code"`
	IsAuthorized bool        `json:"is_authorized"`
	RedirectUrl  string      `json:"redirect_url"`
}

func (s SessionController) GetSessionUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	session := models.Session{}
	if _, err := session.GetById(uint(idInt), s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	// Session not connected
	if !session.Connected {
		c.JSON(http.StatusNoContent, nil)
		return
	}
	user := models.User{}
	if _, err := user.GetById(session.UserId, s.db); err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			user.Id = session.UserId
			if err := user.Save(s.db); err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"errr": err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"errr": err.Error()})
			return
		}
	}
	r := SessionResponse{}
	r.User = user

	if user.IsAccepted {
		r.IsAuthorized = true
		r.Code = uuid.NewString()
		r.RedirectUrl = session.RedirectUrl
		session.Code = r.Code
		if err := session.Save(s.db); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"errr": err.Error()})
			return
		}

	}

	c.JSON(http.StatusOK, r)

}

type ExchangeTokenRequest struct {
	Code         string `json:"code"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (s SessionController) ExchangeToken(c *gin.Context) {

	body := ExchangeTokenRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"err": err.Error()})
		return
	}
	session := models.Session{}
	if _, err := session.GetByCode(body.Code, s.db); err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"err": err.Error()})
		return
	}
	user := models.User{}
	if _, err := user.GetById(session.UserId, s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"err": err.Error()})
		return
	}

	client := models.Client{}
	if _, err := client.GetById(session.ClientId, s.db); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"err": err.Error()})
		return
	}
	if client.Secret != body.ClientSecret {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"err": "Bad client secret"})
		return
	}

	jwt, err := utils.CreateAccessToken(1*time.Hour, user, client.PrivateKey, client.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"acces_token": jwt})

}
