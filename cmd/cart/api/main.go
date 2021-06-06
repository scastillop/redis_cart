package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/scastillop/redis_cart/internal/api/routes"
	"github.com/go-redis/redis"
)

func init() {
}

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	var port string
	app := fiber.New(fiber.Config{
		ReadBufferSize: 9000,
	})

	routes.SetupRoutes(app)

	if port = os.Getenv("HTTP_PORT"); port == "" {
		port = "8080"
	}

	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		Password: "",
		DB: 0,
	})

	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
    if err != nil {
        fmt.Println(err)
    }

    err = client.Set("id1234", json, 0).Err()
    if err != nil {
        fmt.Println(err)
    }
    val, err := client.Get("id1234").Result()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(val)

	log.Fatal(app.Listen(":" + port))
}
