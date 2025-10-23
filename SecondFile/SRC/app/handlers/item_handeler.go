package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/you/fiber-crud/SRC/app/module"
	"github.com/you/fiber-crud/SRC/app/utils"
)

var items = []module.Items{}

var idCount = 1

// create entry
func CreateItem(c *fiber.Ctx) error {

	item := new(module.Items)
	fmt.Println(string(c.BodyRaw()))
	if err := c.BodyParser(item); err != nil {
		fmt.Println("error : -", err)
		return utils.ErrorResponce(c, fiber.StatusBadRequest, "invalid JSON payload")

	}

	if item.Name == "" || item.Price <= 0 {
		return utils.ErrorResponce(c, fiber.StatusBadRequest, "Name and price are required")
	}

	item.ID = idCount
	idCount++
	items = append(items, *item)

	return utils.JSONResponce(c, fiber.StatusCreated, items, "item create success fully")

}

// get all items

func GetAllItems(c *fiber.Ctx) error {
	return utils.JSONResponce(c, fiber.StatusOK, items, "your all user success fully fetch")
}

// get single item
func GetItemById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponce(c, fiber.StatusBadRequest, "invalid Id")
	}
	for _, item := range items {
		if item.ID == id {
			return utils.JSONResponce(c, fiber.StatusOK, item, "item fetch success fully")
		}
	}
	return utils.ErrorResponce(c, fiber.StatusNotFound, "item not found")
}

// update item
func Update_items_By_ID(c *fiber.Ctx) error {

	id, error := strconv.Atoi(c.Params("id"))

	if error != nil {
		return utils.ErrorResponce(c, fiber.StatusBadRequest, "invalid user id")
	}

	var updated module.Items
	userbody := c.Body()
	fmt.Println(userbody)
	if err := c.BodyParser(&updated); err != nil {
		return utils.ErrorResponce(c, fiber.StatusBadRequest, "user payload messing")
	}

	for i, item := range items {
		if item.ID == id {
			if updated.Name != "" {
				items[i].Name = updated.Name
			}
			if updated.Price > 0 {
				items[i].Price = updated.Price
			}
			return utils.JSONResponce(c, fiber.StatusOK, items[i], "success fully  update data")
		}

	}

	return utils.ErrorResponce(c, fiber.StatusNotFound, "items not found")
}

// delete items

func DeleteItems(c *fiber.Ctx) error {

	id, error := strconv.Atoi(c.Params("id"))
	if error != nil {
		return utils.ErrorResponce(c, fiber.StatusBadRequest, "invalid user id")
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return utils.JSONResponce(c, fiber.StatusOK, nil, "item delete success fully")
		}
	}
	return utils.ErrorResponce(c, fiber.StatusNotFound, "items not found")
}
