package controller

import (
	"api-perpus/database"
	"api-perpus/model"
	"api-perpus/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllBooks(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		book := []model.Book{}

		if err := database.GetAll(db, &book); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusNotFound, err, nil)
		}
		return util.ResponseHTTP(ctx, fiber.StatusOK, nil, book)
	}
}

func GetBook(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		book := model.Book{}

		if err := database.GetWhere(db, &book, "id = "+ctx.Params("id")); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusNotFound, err, nil)
		}
		return util.ResponseHTTP(ctx, fiber.StatusOK, nil, book)
	}
}

func SaveBook(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		book := model.Book{}

		if err := ctx.BodyParser(&book); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusBadRequest, err, nil)
		}

		if err := database.SaveData(db, &book); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}

		return util.ResponseHTTP(ctx, fiber.StatusOK, nil, book)
	}
}

func UpdateBook(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		book := model.Book{}

		if err := ctx.BodyParser(&book); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusBadRequest, err, nil)
		}

		if err := database.UpdateData(db, "id = "+ctx.Params("id"), &book); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}

		return util.ResponseHTTP(ctx, fiber.StatusOK, nil, book)
	}
}

func DeleteBook(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		if err := database.DeleteData(db, &model.Book{}, ctx.Params("id")); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}

		return util.ResponseHTTP(ctx, fiber.StatusOK, nil, nil)
	}
}