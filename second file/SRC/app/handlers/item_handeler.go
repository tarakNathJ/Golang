package handlers

import (
	"SRC/app/utils"
	"SecondFile/SRC/app/module"
	"strconv"

	"github.com/gofiber/fiber/v2"
	// "github.com/you/fiber-crud/SRC/app/module"
	// "github.com/you/fiber-crud/SRC/app/utils"
)

var items = []module.Items{}

var idCount = 1

// create entry
func CreateItem(c *fiber.Ctx) error {
	items := new(module.Item)
	if err := c.BodyParser(items); err != nil {
		return utils.ErrorResponce(c, fiber.StatusBadRequest, "invalid JSON payload")

	}
	if items.Name == "" || items.Price <= 0 {
		return utils.ErrorResponce(c, fiber.StatusBadRequest, "Name and price are required")
	}

	items.ID = idCount
	idCount++
	items = append(items, *items)

	return utils.JSONResponce(c, fiber.StatusCreated, items, "item create success fully")

}

// get all entrys

func GetItemsByID(c *fiber.Ctx) error {

	id, error := strconv.Atoi(c.Params("id"))
	if error != nil {
		return utils.E
	}

}
