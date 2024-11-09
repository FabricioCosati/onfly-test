package main

import (
	"log"
	"os"

	"github.com/FabricioCosati/onfly-test/internal/config"
	"github.com/FabricioCosati/onfly-test/internal/di"
	"github.com/FabricioCosati/onfly-test/internal/routes"
	_ "github.com/joho/godotenv"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("error on init environment config: %s", err)
		return
	}

	app, err := di.InitApplication(config)
	if err != nil {
		log.Fatalf("error on init application: %s", err)
		return
	}

	routes.InitOrderRoutes(app)

	port := os.Getenv("SERVER_PORT")
	app.Server.StartServer(port)

}
