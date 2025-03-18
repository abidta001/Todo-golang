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

	//Univeral
	app.Get("/listusers", controllers.ListUser)
	app.Get("/viewtasks", controllers.ViewTasks)

	//User Routes
	user := app.Group("/user")
	user.Post("/signup", controllers.SignupUser)
	user.Post("/login", controllers.LoginUser)

	//Task Routes
	task := app.Group("/task")
	task.Post("/create", middleware.JWTMiddleware(), controllers.CreateTask)

	app.Listen(":3000")
}
