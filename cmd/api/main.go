package main

import (
	"log"
	"test-task-go/docs"
	"test-task-go/internal/config"
	"test-task-go/internal/controller"
	"test-task-go/internal/db"
	"test-task-go/internal/middleware"
	"test-task-go/internal/repository"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title test-task-go
// @host localhost:8080
// @BasePath  /
// @Security ApiKeyAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Get()
	db := db.Get(cfg)

	serviceRepo := repository.NewServiceRepository(db)
	userRepo := repository.NewUserRepository(db)

	serviceController := controller.NewServiceController(serviceRepo)

	router := gin.Default()
	serviceRouter := router.Group("services").Use(middleware.AuthMiddleware(userRepo))
	serviceRouter.GET("", serviceController.Index)
	serviceRouter.GET("/:id", serviceController.Get)
	serviceRouter.POST("", serviceController.Create)
	serviceRouter.DELETE("/:id", serviceController.Delete)

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":" + cfg.Port)

	return err
}
