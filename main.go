package main

import (
	"fiber_demo/model"
	"fmt"
	"github.com/CoderVlogger/go-web-frameworks/pkg"
	"github.com/google/uuid"
	"strconv"

	"github.com/gofiber/fiber/v2"

	_ "github.com/CoderVlogger/go-web-frameworks/pkg"

	"fiber_demo/db"
	_ "fiber_demo/model"
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

	db.ConnectDB()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(JSONTextResponse{Message: "Hello, Fiber"})
	})

	entitiesAPI := app.Group("/entities")
	entitiesAPI.Get("/", entitiesList)
	entitiesAPI.Get("/:id", entitiesGet)
	entitiesAPI.Post("/", entitiesPost)
	entitiesAPI.Put("/", entitiesUpdate)
	entitiesAPI.Delete("/:id", entitiesDelete)
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

	db := db.DB
	var notes []model.Note

	db.Find(&notes)
	if len(notes) == 0 {
		// 없는페이지면 에러메시지와 404로 응답
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}
	// Else return notes(pagination: 4 elem for 1 page
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes[(page-1)*4 : (page*4)-1]})
}

func entitiesGet(c *fiber.Ctx) error {
	// paging
	entityId := c.Params("id")

	entity, err := entitiesRepo.Get(entityId)
	// 없는페이지면 에러메시지와 404로 응답
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusNotFound).JSON(errMsg)
	}
	return c.JSON(entity)
}

func entitiesPost(c *fiber.Ctx) error {

	db := db.DB
	note := new(model.Note)

	err := c.BodyParser(note)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	note.ID = uuid.New()
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

func entitiesUpdate(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = entitiesRepo.Update(&entity)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	return c.JSON(entity)

}

func entitiesDelete(c *fiber.Ctx) error {

	entityId := c.Params("id")

	err := entitiesRepo.Delete(entityId)
	// 없는페이지면 에러메시지와 404로 응답
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}
	return c.JSON(pkg.TextResponse{Message: "entity deleted"})
}
