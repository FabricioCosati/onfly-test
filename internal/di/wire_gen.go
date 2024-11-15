// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/FabricioCosati/onfly-test/internal/api"
	"github.com/FabricioCosati/onfly-test/internal/config"
	"github.com/FabricioCosati/onfly-test/internal/controllers"
	"github.com/FabricioCosati/onfly-test/internal/database"
	"github.com/FabricioCosati/onfly-test/internal/repository"
	"github.com/FabricioCosati/onfly-test/internal/services"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitApplication(cfg config.Config) (*AppInit, error) {
	gormDB, err := db.ConnectToDatabase(cfg)
	if err != nil {
		return nil, err
	}
	orderRepositoryImpl := repository.OrderRepositoryInit(gormDB)
	orderServiceImpl := services.OrderServiceInit(orderRepositoryImpl)
	orderControllerImpl := controllers.OrderControllerInit(orderServiceImpl)
	orderInit := NewOrderInitialization(orderRepositoryImpl, orderServiceImpl, orderControllerImpl)
	appInit := NewAppInit(orderInit, gormDB)
	return appInit, nil
}

// wire.go:

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
