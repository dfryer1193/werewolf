package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dfryer1193/werewolf/internal/api"
	"github.com/dfryer1193/werewolf/internal/config"
	"github.com/dfryer1193/werewolf/internal/db"
	"github.com/dfryer1193/werewolf/internal/logging"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.LoadConfig()

	logging.InitLogger(cfg.LogLevel)

	pgdb, err := db.NewPostgresDB(cfg.DBConnString)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to database")
	}
	defer pgdb.Close()

	r := mux.NewRouter()
	werewolfApi := api.NewAPI()
	werewolfApi.RegisterRoutes(r)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	log.Info().Msg("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited properly")
}
