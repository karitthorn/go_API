package handlers

import (
    "time"
    "github.com/karitthorn/go_API/models"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)


func GetTodos(c *fiber.Ctx) error {
    return c.JSON(models.Todos)
}


func GetTodo(c *fiber.Ctx) error {
    id := c.Params("id")

    for _, todo := range models.Todos {
        if todo.ID == id {
            return c.JSON(todo)
        }
    }

    return c.Status(404).JSON(fiber.Map{
        "error": "Todo not found",
    })
}


func CreateTodo(c *fiber.Ctx) error {
    todo := new(models.Todo)

    if err := c.BodyParser(todo); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Validate
    if todo.Title == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "Title is required",
        })
    }

    todo.ID = uuid.New().String()
    todo.CreatedAt = time.Now()
    todo.Completed = false

    models.Todos = append(models.Todos, *todo)

    return c.Status(201).JSON(todo)
}


func UpdateTodo(c *fiber.Ctx) error {
    id := c.Params("id")

    type UpdateInput struct {
        Title       *string `json:"title"`
        Description *string `json:"description"`
        Completed   *bool   `json:"completed"`
    }

    input := new(UpdateInput)
    if err := c.BodyParser(input); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    for i, todo := range models.Todos {
        if todo.ID == id {
            if input.Title != nil {
                models.Todos[i].Title = *input.Title
            }
            if input.Description != nil {
                models.Todos[i].Description = *input.Description
            }
            if input.Completed != nil {
                models.Todos[i].Completed = *input.Completed
            }

            return c.JSON(models.Todos[i])
        }
    }

    return c.Status(404).JSON(fiber.Map{
        "error": "Todo not found",
    })
}


func DeleteTodo(c *fiber.Ctx) error {
    id := c.Params("id")

    for i, todo := range models.Todos {
        if todo.ID == id {
            models.Todos = append(models.Todos[:i], models.Todos[i+1:]...)
            return c.SendStatus(204)
        }
    }

    return c.Status(404).JSON(fiber.Map{
        "error": "Todo not found here",
    })
}

func GetHealth(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "status": "OK",
        "time":   time.Now(),
    })
}

