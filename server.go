package main

import (
	"github.com/gofiber/fiber/v2"
)

type Handler = func(c *fiber.Ctx) error
type FiberContext = *fiber.Ctx

func getTests(c FiberContext) error {
	return c.SendString("This will get all tests from a queue!")
}

func getTestById(c FiberContext) error {
	return c.SendString("This will get a test by the ID in the route: " + c.Params("id") + "!")
}

func runTest(c *fiber.Ctx) error {
	return c.SendString("This will add a test to a queue, and return a job id!")
}

func main() {
	app := fiber.New()

	app.Get("/ping", func(c FiberContext) error {
		if c.Query("name") != "" {
			return c.SendString("pong, " + c.Query("name") + "!")
		}
		return c.SendString("pong!")
	})

	tests := app.Group("/tests")
	tests.Get("/", getTests)
	tests.Get("/:id", getTestById)
	tests.Post("/run", runTest)

	app.Listen(":9090")
}
