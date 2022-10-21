package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"belajar/tugas1/controllers"
)

func main() {
	// load template engine
	engine := html.New("./views",".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public","./public")

	// controllers
	helloController := controllers.InitHelloController()
	prodController := controllers.InitProductController()

	p := app.Group("/greetings")
	p.Get("/", helloController.Greeting)
	p.Get("/hello", helloController.SayHello)
	p.Get("/myview", helloController.HelloView)

	prod := app.Group("/products")
	prod.Get("/", prodController.IndexProduct)
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Get("/productdetail", prodController.GetDetailProduct)
	prod.Get("/detail/:id", prodController.GetDetailProduct2)

	app.Listen(":3000")
}