package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SladeThe/word-of-wisdom/internal/client"
	"github.com/SladeThe/word-of-wisdom/internal/client/config"
)

func main() {
	ctx := context.Background()

	log.Print("[INFO] initialize config")
	cfg, errCfg := config.New()
	if errCfg != nil {
		log.Fatal("failed initializing config: ", errCfg)
	}

	log.Printf("[INFO] connect to %s:%d", cfg.Client.Host, cfg.Client.Port)
	cln, errClient := client.Start(ctx, cfg.Client)
	if errClient != nil {
		log.Fatal("failed dialing TCP server: ", errClient)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[INFO] shutdown gracefully")

	cln.Shutdown()

	log.Println("[INFO] bye")
}
