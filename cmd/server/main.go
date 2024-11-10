package main

import (
	"context"
	"log"
	"os"

	"github.com/FabricioCosati/onfly-test/internal/config"
	"github.com/FabricioCosati/onfly-test/internal/di"
	"github.com/FabricioCosati/onfly-test/internal/middlewares"
	"github.com/FabricioCosati/onfly-test/internal/routes"
	_ "github.com/joho/godotenv"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("error on init environment config: %s", err)
	}

	app, err := di.InitApplication(config)
	if err != nil {
		log.Fatalf("error on init application: %s", err)
	}

	tp, err := middlewares.InitTracerMetrics()
	if err != nil {
		log.Fatalf("error on init application: %s", err)
	}
	defer tp.Shutdown(context.Background())

	routes.InitOrderRoutes(app)

	port := os.Getenv("SERVER_PORT")
	app.Server.StartServer(port)
}
