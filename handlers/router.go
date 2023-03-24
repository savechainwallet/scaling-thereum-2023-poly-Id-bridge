package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/config"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, conf *config.AppConfig) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	clientsService := Clients{db: db}

	polyIdService := PolygonIDService{
		RedirectUrl:        fmt.Sprintf("%s/callback", conf.AppDomain),
		RPCUrl:             conf.RPCUrl,
		ValidatorsContract: conf.ValidatorContract,
		ResolverPrefix:     conf.ResolverPrefix,
		KeyDir:             "./keys",
		db:                 db,
	}
	sessionService := SessionController{
		db: db,
	}

	router.POST("/clients", clientsService.Register)
	router.GET("/api/qr", polyIdService.GetBasicQR)
	router.POST("/api/sessions/:id", polyIdService.GenerateProof)

	router.POST("/callback", polyIdService.ServeRedirect)
	router.GET("/api/sessions/:id", sessionService.GetSessionUser)

	return router

}
