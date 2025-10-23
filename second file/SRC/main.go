package main

// "strconv"

// "github.com/gofiber/fiber/v2"
// "github.com/gofiber/fiber/v2/middleware/logger"

import (
	"fmt"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []user

func main() {

	// app := fiber.New()

	// app.Use(logger.New())

	// app.Listen(":3000")

	items1 := user{ID: 1, Name: "tarak", Price: 2309}

	items = append(items, items1)
	items = append(items, user{ID: 2, Name: "TarakJana", Price: 6000})
	items3 := user{ID: 3, Name: "gulo", Price: 20989}
	index := 1

	items = append(items[:index], append([]user{items3}, items[index:]...)...)

	fmt.Println(items)

	items = append(items[:index], items[index+1:]...)
	fmt.Println(items)

	items[1].Name = "chodu"
	fmt.Println(items)

}
