package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/shankar7042/hotel-reservation-golang/api"
	"github.com/shankar7042/hotel-reservation-golang/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"

var config = fiber.Config{
	ErrorHandler: func(c fiber.Ctx, err error) error {
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the api server")
	flag.Parse()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)
}
