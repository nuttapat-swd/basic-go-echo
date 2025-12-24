package main

import (
	"go_poc/internal/school"
	"go_poc/pkg/database"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Instance Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	schoolRepo := school.NewSchoolRepository(db)
	schoolService := school.NewSchoolService(schoolRepo)
	schoolHandler := school.NewSchoolHandler(schoolService)

	// Routes
	api_v1 := e.Group("api/v1")

	api_v1.GET("/schools", schoolHandler.List)
	api_v1.GET("/schools/:id", schoolHandler.Get)
	api_v1.POST("/schools", schoolHandler.Create)
	api_v1.PUT("/schools/:id", schoolHandler.Update)
	api_v1.DELETE("/schools/:id", schoolHandler.Delete)

	// Server
	e.Logger.Fatal(e.Start(":8080"))
}
