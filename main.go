package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/8BITS-COLAB/nfebot/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	v1 := app.Group("/api/v1")

	bot := nfebot.New()

	v1.Post("/", func(c *fiber.Ctx) error {
		var issueNFEDTO nfebot.IssueNFEDTO

		if err := c.BodyParser(&issueNFEDTO); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		image, err := bot.WithRetries(3, issueNFEDTO)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		c.Response().Header.SetContentType("image/png")

		filename := fmt.Sprintf("data/%d.png", time.Now().Unix())
		ioutil.WriteFile(filename, image, 0644)

		return c.Status(http.StatusOK).Send(image)
	})

	log.Fatal(app.Listen(":4000"))
}
