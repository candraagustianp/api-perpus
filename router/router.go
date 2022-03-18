package router

import (
	"api-perpus/controller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetRouting(app *fiber.App, db *gorm.DB) {
	router := app.Group("/api")
	PublicRouting(router, db)
}

func PublicRouting(router fiber.Router, db *gorm.DB) {
	router.Get("/book", controller.GetAllBooks(db))

	router.Get("/book/:id", controller.GetBook(db))

	router.Post("/book", controller.SaveBook(db))

	router.Put("/book/:id", controller.UpdateBook(db))

	router.Delete("/book/:id", controller.DeleteBook(db))

}
