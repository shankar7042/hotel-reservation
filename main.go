package main

import (
	"flag"

	"github.com/gofiber/fiber/v3"
	"github.com/shankar7042/hotel-reservation-golang/api"
)

func main() {

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the api server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
