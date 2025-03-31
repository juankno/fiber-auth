package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juankno/fiber-auth/database"
	"github.com/juankno/fiber-auth/model"
)

type bookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func GetBooks(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "All books",
	})
}

func GetBook(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Get book",
	})
}

func CreateBook(c *fiber.Ctx) error {
	var req bookRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	book := model.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	res := database.DB.Create(&book)

	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Book created",
	})
}

func UpdateBook(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Book created",
	})
}

func DeleteBook(c *fiber.Ctx) error {
	return c.Status(204).JSON(fiber.Map{})
}
