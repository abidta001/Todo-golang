package controllers

import (
	"todo/database"
	"todo/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	err := c.BodyParser(&task)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating Task"})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}
func ViewTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch data")
	}
	return c.Status(fiber.StatusOK).JSON(tasks)
}
