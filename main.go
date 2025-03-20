package main

import (
	"todo/controllers"
	"todo/database"
	"todo/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDB()

	//User Routes
	user := app.Group("/user")
	user.Post("/signup", controllers.SignupUser)
	user.Post("/login", controllers.LoginUser)

	//Task Routes
	task := app.Group("/task")
	task.Post("/create", middleware.JWTMiddleware(), controllers.CreateTask)
	task.Get("/view", middleware.JWTMiddleware(), controllers.ViewTasks)
	task.Put("/update/:id", middleware.JWTMiddleware(), controllers.UpdateTask)

	app.Listen(":3000")
}
