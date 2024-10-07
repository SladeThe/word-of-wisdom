package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SladeThe/word-of-wisdom/internal/server"
	"github.com/SladeThe/word-of-wisdom/internal/server/config"
	"github.com/SladeThe/word-of-wisdom/internal/server/repositories"
	embeddedRepositories "github.com/SladeThe/word-of-wisdom/internal/server/repositories/embedded"
	"github.com/SladeThe/word-of-wisdom/internal/server/services"
	implServices "github.com/SladeThe/word-of-wisdom/internal/server/services/impl"
)

func main() {
	ctx := context.Background()

	log.Print("[INFO] initialize config")
	cfg, errCfg := config.New()
	if errCfg != nil {
		log.Fatal("[ERROR] failed initializing config: ", errCfg)
	}
	ctx = config.Set(ctx, cfg)

	log.Print("[INFO] initialize repositories")
	rr, errRepositories := initRepositories()
	if errRepositories != nil {
		log.Fatal("[ERROR] failed initializing repositories: ", errRepositories)
	}
	ctx = repositories.Set(ctx, rr)

	log.Print("[INFO] initialize services")
	ss, errServices := initServices(cfg.Services)
	if errServices != nil {
		log.Fatal("[ERROR] failed initializing services: ", errServices)
	}
	ctx = services.Set(ctx, ss)

	log.Printf("[INFO] start server on :%d", cfg.Server.Port)
	srv, errServer := server.Start(ctx, cfg.Server)
	if errServer != nil {
		log.Fatal("[ERROR] failed starting TCP server: ", errServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[INFO] shutdown gracefully")

	srv.Shutdown()

	log.Println("[INFO] bye")
}

func initRepositories() (repositories.Repositories, error) {
	return repositories.New(embeddedRepositories.NewWordOfWisdom()), nil
}

func initServices(cfg config.Services) (services.Services, error) {
	challenge, errChallenge := implServices.NewChallenge(cfg.Challenge)
	if errChallenge != nil {
		return services.Services{}, errChallenge
	}

	return services.New(
		challenge,
		implServices.NewWordOfWisdom(),
	), nil
}
