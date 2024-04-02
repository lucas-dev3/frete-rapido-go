package main

import (
	"context"
	"log"

	"github.com/lucas-dev3/frete-rapido-go.git/config"
	"github.com/lucas-dev3/frete-rapido-go.git/database"
	"github.com/lucas-dev3/frete-rapido-go.git/internal/http/gin"
	"github.com/lucas-dev3/frete-rapido-go.git/internal/http/webserver"
	"github.com/lucas-dev3/frete-rapido-go.git/quote"
	Repositories "github.com/lucas-dev3/frete-rapido-go.git/quote/postgres"
)

func main() {
	db := database.Connection()
	defer db.Close(context.Background())

	envs := config.LoadEnvVars()

	quoteRepo := Repositories.NewQuoteRepository(db)
	quoteService := quote.NewService(quoteRepo)

	h := gin.Handlers(envs, quoteService)

	if err := webserver.Start(envs.APIPort, h); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
