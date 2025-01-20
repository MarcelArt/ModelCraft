package view_handlers

import (
	"github.com/MarcelArt/app_standard/utils"
	"github.com/MarcelArt/app_standard/views/hello"
	"github.com/gofiber/fiber/v2"
)

func HelloWorldView(c *fiber.Ctx) error {
	return utils.Render(c, hello.Show("marcel"))
}
