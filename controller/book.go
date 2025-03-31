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

	var books []model.Book

	if err := database.DB.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving books",
		})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var req bookRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	book := model.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	res := database.DB.Create(&book)

	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": res.Error.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created",
	})
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var req bookRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var book model.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	book.Title = req.Title
	book.Author = req.Author

	if err := database.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating book",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book updated",
	})
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&model.Book{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting book",
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
