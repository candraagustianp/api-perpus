package controller

import (
	"api-perpus/database"
	"api-perpus/model"
	"api-perpus/util"
	"fmt"

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

		//upload file
		fileName := fmt.Sprintf("%s_%s.pdf", book.Isbn, book.Judul)
		fileDecode64, err := util.Base64Decode(book.File)
		if err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}
		if err := util.WriteFileBase64("files", fileName, fileDecode64); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}
		book.File = fileName

		if err := database.SaveData(db, &book); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}

		return util.ResponseHTTP(ctx, fiber.StatusOK, nil, book)
	}
}

func UpdateBook(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		newBook := model.Book{}
		oldBook := model.Book{}

		if err := ctx.BodyParser(&newBook); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusBadRequest, err, nil)
		}
		if err := database.GetWhere(db, &oldBook, "id = "+ctx.Params("id")); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}

		fileName := fmt.Sprintf("%s_%s.pdf", newBook.Isbn, newBook.Judul)

		if util.IsBase64(newBook.File) {
			if e := util.RemoveFile("files", oldBook.File); e != nil {
				return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, e, nil)
			}

			fileDecode64, err := util.Base64Decode(newBook.File)
			if err != nil {
				return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
			}
			if err := util.WriteFileBase64("files", fileName, fileDecode64); err != nil {
				return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
			}

		}

		if fileName != oldBook.File {
			if e := util.Rename("files", oldBook.File, fileName); e != nil {
				return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, e, nil)
			}
		}

		newBook.File = fileName

		if err := database.UpdateData(db, "id = "+ctx.Params("id"), &newBook); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}

		return util.ResponseHTTP(ctx, fiber.StatusOK, nil, newBook)
	}
}

func DeleteBook(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		book := model.Book{}

		if err := database.GetWhere(db, &book, "id = "+ctx.Params("id")); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusNotFound, err, nil)
		}
		if e := util.RemoveFile("files", book.File); e != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, e, nil)
		}

		if err := database.DeleteData(db, &book, ctx.Params("id")); err != nil {
			return util.ResponseHTTP(ctx, fiber.StatusInternalServerError, err, nil)
		}

		return util.ResponseHTTP(ctx, fiber.StatusOK, nil, nil)
	}
}
