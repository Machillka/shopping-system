package main

import (
	"log"

	httpadapter "github.com/machillka/shopping-system/internal/adapters/http"

	"github.com/gin-gonic/gin"
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
	domainSvc := domain.DefalultOrderDomainService{}
	// 依赖注入
	orderSvc := application.NewOrderService(orderRepo, domainSvc)

	gin.SetMode(gin.ReleaseMode)
	rounter := gin.New()
	rounter.Use(gin.Logger(), gin.Recovery())

	handler := httpadapter.NewOrderHandler(orderSvc)
	handler.RegisterRoutes(rounter)

	addr := viper.GetString("server.address")
	log.Printf("Starting server in %s", addr)
	if err := rounter.Run(addr); err != nil {
		log.Fatal("server error:", err)
	}

}
