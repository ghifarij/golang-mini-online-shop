package main

import (
	"log"

	"github.com/ghifarij/golang-mini-online-shop/apps/auth"
	"github.com/ghifarij/golang-mini-online-shop/external/database"
	"github.com/ghifarij/golang-mini-online-shop/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
