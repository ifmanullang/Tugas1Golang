package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HelloController struct {
	// declare variables

}
func InitHelloController() *HelloController {
	return &HelloController{}
}

func (controller *HelloController) Greeting(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "welcome...",
	})
}
func (controller *HelloController) SayHello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "say hello...",
	})
}
func (controller *HelloController) HelloView(c *fiber.Ctx) error {
	return c.Render("myview", fiber.Map{
		"Title": "ini judul...",
	})
}
