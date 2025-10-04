package main

import (
	"log"
	"net/http"

	"github.com/codeid/hr-api-simple/internal/handlers"
	"github.com/codeid/hr-api-simple/internal/models"
	"github.com/codeid/hr-api-simple/internal/repositories"
	"github.com/codeid/hr-api-simple/internal/services"
	"github.com/codeid/hr-api-simple/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//1. initialize database
	db, err := database.NewDB()
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	//2. create schema hr
	db.Exec(("CREATE SCHEMA IF NOT EXSITS hr"))

	//3. AutoMigrate
	err = db.AutoMigrate(&models.Region{})
	if err != nil {
		log.Fatal("Failed to migrate tables : ", err)
	}

	//4.initial repository
	regionRepo := repositories.NewRegionRepository(db)

	//5.init service
	regionService := services.NewRegionService(regionRepo)

	//6.init handler
	regionHandler := handlers.NewRegionHandler(regionService)

	//setup routes
	router := gin.Default()

	//call handler

	api := router.Group("/api")
	{
		router.GET("/", welcomeHandler)

		//routers endpoint
		regions := api.Group("/regions")
		{
			regions.GET("", regionHandler.GetRegions)
			regions.GET("/:id", regionHandler.GetRegion)
			regions.POST("", regionHandler.CreateRegion)
			regions.PUT("/:id", regionHandler.UpdateRegion)
			regions.DELETE("/:id", regionHandler.DeleteRegion)
		}
	}

	log.Println("server starting on :8080")
	router.Run(":8080")

}

func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to gin framework",
		"status":  "running",
	})
}
