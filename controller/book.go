package controller

import "github.com/gofiber/fiber/v2"

type bookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func GetBooks(c *fiber.Ctx) error {
	//
}

func GetBook(c *fiber.Ctx) error {
	//
}

func CreateBook(c *fiber.Ctx) error {
	//
}

func UpdateBook(c *fiber.Ctx) error {
	//
}

func DeleteBook(c *fiber.Ctx) error {
	//
}
