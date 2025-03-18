package main

import (
	"todo/controllers"
	"todo/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDB()
	user := app.Group("/user")
	user.Post("/signup", controllers.SignupUser)
	user.Post("/login", controllers.LoginUser)
	user.Get("/", controllers.ListUser)

	task := app.Group("/task")
	task.Post("/create", controllers.CreateTask)
	task.Get("/", controllers.ViewTasks)
	app.Listen(":3000")
}
