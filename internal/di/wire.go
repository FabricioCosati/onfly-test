//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FabricioCosati/onfly-test/internal/api"
	"github.com/FabricioCosati/onfly-test/internal/config"
	db "github.com/FabricioCosati/onfly-test/internal/database"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type AppInit struct {
	OrderInit *OrderInit
	Server    *api.Server
	Database  *gorm.DB
}

func NewAppInit(
	order *OrderInit,
	database *gorm.DB,
) *AppInit {
	return &AppInit{
		OrderInit: order,
		Database:  database,
		Server:    api.InitServer(),
	}
}

func InitApplication(cfg config.Config) (*AppInit, error) {
	wire.Build(
		OrderProviders,
		db.ConnectToDatabase,
		NewAppInit)
	return &AppInit{}, nil
}
