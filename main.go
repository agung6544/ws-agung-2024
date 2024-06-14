package main

import (
	"log"

	"github.com/agung6544/ws-agung-2024/config"
	_ "github.com/agung6544/ws-agung-2024/docs"
	"github.com/agung6544/ws-agung-2024/url"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/agung6544
// @contact.email 714220039@std.ulbi.ac.id

// @host https://ws-agung-2024-b77249100221.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
