package main

import (
	"log"
	// "fmt"
	"todo/config"
	"todo/connectors"
	"todo/repos"
	"todo/routes"
	"todo/services"
	// "todo/connectors"
	"github.com/gin-gonic/gin"
	"todo/models"
	// "todo/controller"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	// gosql "github.com/go-sql-driver/mysql"
)

func main() {

	cfg := config.AppConfigF()

	log.Println("Starting server...")

	dbs, err := database.New(cfg)

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	router := gin.Default() //Init Router

	appConfig := config.AppConfig{
		Router: router,
		Dbs:    dbs,
		Cfg:    cfg,
	}

	todoRepo := repos.TodoRepo{
		DB: appConfig.Dbs.DB,
	}

	models.ConnectDatabase()

	TodoService := services.NewService(services.ServiceConfig{
		TodoRepo: todoRepo,
	})

	routes.InitRoutes(config.HandlerConfig{
		R:           appConfig.Router,
		TodoService: TodoService,
	})

	router.Run(":" + cfg.Server.Port)

}
