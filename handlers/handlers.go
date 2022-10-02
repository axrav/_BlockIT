package handlers

import (
	"blockit/utils"

	"github.com/gofiber/fiber/v2"
)

func NewData(c *fiber.Ctx) error {
	//Data := new(utils.SampleData)
	var Data *utils.SampleData
	err := c.BodyParser(&Data)

	if err != nil {
		return c.JSON(fiber.Map{"error": err})
	}
	block := *utils.CurrentBlock()
	block.AddBlock(Data)
	return c.Status(200).JSON(Data)
}
func GetBlock(c *fiber.Ctx) error {
	block := *utils.CurrentBlock()
	return c.Status(200).JSON(fiber.Map{"data": block.Blocks})
}
