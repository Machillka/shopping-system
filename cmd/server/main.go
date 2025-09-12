package main

import (
	"log"

	"github.com/machillka/shopping-system/internal/application"
	"github.com/machillka/shopping-system/internal/domain"
	"github.com/machillka/shopping-system/internal/infra/sqlite"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	path := viper.GetString("database.sqlite.path")
	if err := sqlite.Init(path); err != nil {
		log.Fatal(err)
	}
	defer sqlite.Close()

	orderRepo := sqlite.NewOrderRepository()
	domainSvc := domain.DefalultOrderDomainService()
	// 依赖注入
	orderSvc := application.NewOrderService(orderRepo, domainSvc)
	
}