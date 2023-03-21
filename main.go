package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/config"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/handlers"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func main() {
	conf := config.GetConfig()
	ctx, cancel := context.WithCancel(context.Background())

	r := handlers.NewRouter()

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
