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
func UpdateTask(c *fiber.Ctx) error {
	taskID := c.Params("id")

	var input struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}
	err := c.BodyParser(&input)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request"})
	}
	user_id, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid User"})
	}
	var task models.Task
	if err := database.DB.Where("id=? and user_id=?", taskID, user_id).First(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Task not found!"})
	}
	task.Title = input.Title
	task.Status = input.Status
	if err := database.DB.Save(&task).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update task"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ID":     task.ID,
		"Title":  task.Title,
		"Status": task.Status,
	})
}
func DeleteTask(c *fiber.Ctx) error {
	taskID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid User"})
	}
	var task models.Task
	if err := database.DB.Where("id=? and user_id=?", taskID, userID).First(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Task not found"})
	}
	if err := database.DB.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete task"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Task deleted successfully"})
}
