package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/config"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/handlers"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/models"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	ctx, cancel := context.WithCancel(context.Background())
	InitLogger(conf.LogLevel)

	var db *gorm.DB
	var err error

	for i := 1; i < 10; i++ {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  conf.DatabaseUrl,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Trying to connect")

		time.Sleep(5 * time.Second)

	}

	if err != nil {
		Logger.Panic("DB INIT ERROR", zap.Error(err))
	}

	db.AutoMigrate(&models.Session{}, &models.Client{}, &models.User{})

	db.Set("gorm:auto_preload", true)

	r := handlers.NewRouter(db, conf)

	srv := &http.Server{
		Addr:    conf.TCPPort,
		Handler: r,
	}

	go func() {
		Logger.Info(fmt.Sprintf("Listen started on port %s", conf.TCPPort))
		if err := srv.ListenAndServe(); err != nil {
			Logger.Panic("Handle server error", zap.Error(err))
		}
	}()

	// Listen for os sygnals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	Logger.Info("App Interrputtes. Waiting for graseful shutdown")
	srv.Shutdown(ctx)
	Logger.Info("Http server stopped")
	cancel()
	os.Exit(0)

}
