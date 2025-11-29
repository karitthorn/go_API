package main

import (
    "log"
    "github.com/karitthorn/go_API/handlers"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New(fiber.Config{
        AppName: "Todo API v1.0.0",
    })

 
    app.Use(logger.New())
    app.Use(cors.New())

 
    api := app.Group("/api/v1")

    api.Get("/todos", handlers.GetTodos)
    api.Get("/health", handlers.GetHealth)
    api.Get("/todos/:id", handlers.GetTodo)
    api.Post("/todos", handlers.CreateTodo)
    api.Put("/todos/:id", handlers.UpdateTodo)
    api.Delete("/todos/:id", handlers.DeleteTodo)

    log.Fatal(app.Listen(":3000"))
}