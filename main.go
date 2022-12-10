package main

import (
	"flag"
	"log"

	"github.com/devhammed/fibery/controllers"
	"github.com/devhammed/fibery/database"
	"github.com/devhammed/fibery/middlewares"
	"github.com/devhammed/fibery/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	flag.Parse()

	database.Connect()

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	app.Route("/api", func(api fiber.Router) {
		api.Use(cors.New())

		api.Route("/v1", func(v1 fiber.Router) {
			v1.Route("/users", func(users fiber.Router) {
				users.Get("/", controllers.ApiV1_GetUsers).Name("api.v1.users.index")
				users.Get("/:id", controllers.ApiV1_GetUser).Name("api.v1.users.show")
				users.Post("/", middlewares.Validator(&requests.CreateUserRequest{}), controllers.ApiV1_CreateUser).Name("api.v1.users.create")
			})
		})
	})

	log.Fatal(app.Listen(*port))
}
