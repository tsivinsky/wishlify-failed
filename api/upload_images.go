package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
)

func HandleUploadImages(c *fiber.Ctx) error {
	fh, err := c.FormFile("file")
	if err != nil {
		return MakeApiError(400, err.Error())
	}

	file, err := fh.Open()
	defer file.Close()

	filePath := fmt.Sprintf("%s/%s", StaticDirectory, fh.Filename)

	savedFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return MakeApiError(500, err.Error())
	}
	defer savedFile.Close()

	_, err = io.Copy(savedFile, file)
	if err != nil {
		return MakeApiError(500, err.Error())
	}

	return c.JSON(fiber.Map{
		"image": filePath,
	})
}
