package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./web", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).Render("errorPage", fiber.Map{
				"Title": "Error",
				"Error": err.Error(),
				"Code":  code,
			})
		},
	})
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Static("/web", "./web")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/dashboard", fiber.Map{
			"Title": "Dashboard",
		}, "index")
	})

	app.Get("/buttons", func(c *fiber.Ctx) error {
		return c.Render("pages/ui-features/buttons", fiber.Map{
			"Title": "Buttons",
		}, "index")
	})

	app.Get("/dropdowns", func(c *fiber.Ctx) error {
		return c.Render("pages/ui-features/dropdowns", fiber.Map{
			"Title": "Dropdowns",
		}, "index")
	})

	app.Get("/typography", func(c *fiber.Ctx) error {
		return c.Render("pages/ui-features/typography", fiber.Map{
			"Title": "Typography",
		}, "index")
	})

	app.Get("/basic-elements", func(c *fiber.Ctx) error {
		return c.Render("pages/forms/basic_elements", fiber.Map{
			"Title": "Basic-elements",
		}, "index")
	})

	app.Get("/chartjs", func(c *fiber.Ctx) error {
		return c.Render("pages/charts/chartjs", fiber.Map{
			"Title": "Chartjs",
		}, "index")
	})

	app.Get("/basic-table", func(c *fiber.Ctx) error {
		return c.Render("pages/tables/basic-table", fiber.Map{
			"Title": "Basic-table",
		}, "index")
	})

	app.Get("/font-awesome", func(c *fiber.Ctx) error {
		return c.Render("pages/icons/font-awesome", fiber.Map{
			"Title": "Font-awesome",
		}, "index")
	})

	app.Get("/blank-page", func(c *fiber.Ctx) error {
		return c.Render("pages/samples/blank-page", fiber.Map{
			"Title": "Blank-page",
		}, "index")
	})

	app.Get("/error-404", func(c *fiber.Ctx) error {
		return c.Render("pages/samples/error-404", fiber.Map{
			"Title": "Error-404",
		})
	})

	app.Get("/error-500", func(c *fiber.Ctx) error {
		return c.Render("pages/samples/error-500", fiber.Map{
			"Title": "Error-500",
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("pages/samples/login", fiber.Map{
			"Title": "Login",
		})
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("pages/samples/register", fiber.Map{
			"Title": "Register",
		})
	})

	app.Get("/documentation", func(c *fiber.Ctx) error {
		return c.Render("docs/documentation", fiber.Map{
			"Title": "Documentation",
		})
	})

	app.Listen(":3000")
}
