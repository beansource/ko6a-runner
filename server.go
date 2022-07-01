package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler = func(c *fiber.Ctx) error
type FiberContext = *fiber.Ctx

type Ko6aTest struct {
	id 					string		// The id of the test run
	scriptUrl 	string 		// The github url of the script to run
	testResult  string 		// Result output of the test
}

func runTest(c *fiber.Ctx) error {
	RunScript(Ko6aTest{
		id: "1",
		scriptUrl: "https://github.com/beansource/stinky-mud-puddle/blob/main/using-k6/03-metrics.js",
	})
	fmt.Println("Ran test")
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
	tests.Get("/run", runTest)

	// go func() {
	// 	ksix.RunScript()
	// }()

	app.Listen(":9090")
}
