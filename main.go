package main

import (
	"fmt"
	"github.com/CoderVlogger/go-web-frameworks/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"

	_ "github.com/CoderVlogger/go-web-frameworks/pkg"
)

type JSONTextResponse struct {
	Message string
}

//interface
var (
	//instance loading data into memory
	entitiesRepo pkg.EntityRepository = pkg.NewEntityMemoryRepository()
)

func main() {
	fmt.Println("hello, world!")

	entitiesRepo.Init()

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(JSONTextResponse{Message: "Hello, Fiber"})
	})

	entitiesAPI := app.Group("/entities")
	entitiesAPI.Get("/", entitiesList)

	app.Listen(":8080")
}

func entitiesList(c *fiber.Ctx) error {
	// paging
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)
	//page 변수를 이상한거 집어넣거나 하면 1페이지로 리턴
	if err != nil {
		page = 1
	}
	entities, err := entitiesRepo.List(page, 4)
	// 없는페이지면 에러메시지와 404로 응답
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusNotFound).JSON(errMsg)
	}
	return c.JSON(entities)
}
