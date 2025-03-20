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
	userID := c.Locals("user_id").(uint)
	task.UserID = userID
	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating Task"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ID":     task.ID,
		"Title":  task.Title,
		"status": task.Status,
	})
}

func ViewTasks(c *fiber.Ctx) error {
	var tasks []struct {
		ID     uint   `json:"id"`
		Title  string `json:"title"`
		Status string `json:"status"`
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user"})
	}

	if err := database.DB.Model(&models.Task{}).
		Select("id, title, status").
		Where("user_id = ?", userID).
		Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch tasks"})
	}

	return c.JSON(tasks)
}
