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
	e.Pre(middleware.RemoveTrailingSlash())

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// School module setup
	schoolRepo := school.NewSchoolRepository(db)
	schoolService := school.NewSchoolService(schoolRepo)
	schoolHandler := school.NewSchoolHandler(schoolService)

	// Classroom module setup
	classroomRepo := school.NewClassroomRepository(db)
	classroomService := school.NewClassroomService(classroomRepo, schoolRepo)
	classroomHandler := school.NewClassroomHandler(classroomService)

	// Routes
	router := e.Group("api/v1")

	router.GET("/schools", schoolHandler.List)
	router.GET("/schools/:id", schoolHandler.Get)
	router.POST("/schools", schoolHandler.Create)
	router.PUT("/schools/:id", schoolHandler.Update)
	router.DELETE("/schools/:id", schoolHandler.Delete)

	router.GET("/classrooms", classroomHandler.List)
	router.GET("/classrooms/:id", classroomHandler.Get)
	router.POST("/classrooms", classroomHandler.Create)
	router.PUT("/classrooms/:id", classroomHandler.Update)
	router.DELETE("/classrooms/:id", classroomHandler.Delete)

	// Server
	e.Logger.Fatal(e.Start(":8080"))
}
