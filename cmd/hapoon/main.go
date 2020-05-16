package main

import (
	"fmt"
	"log"

	"hapoon/api/internal/hapoon/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("main start")
	fmt.Println("Hello, I'm hapoon's api")

	// prepare echo instance
	e := echo.New()

	// prepare middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	profile := e.Group("/profile")
	profileService := service.NewProfiler()
	profile.GET("", profileService.List)
	profile.GET("/:id", profileService.Get)
	profile.POST("", profileService.Create)
	profile.PUT("/:id", profileService.Replace)

	// Start http server
	e.Logger.Fatal(e.Start(":8080"))
}
