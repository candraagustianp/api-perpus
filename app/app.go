package app

import (
	"api-perpus/config"
	"api-perpus/database"
	"api-perpus/model"
	"api-perpus/router"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func Run(config *config.Config) {
	db := connectDatabase(config)
	migrateTabel(db)
	runServer(config, db)
}

func runServer(conf *config.Config, db *gorm.DB) {
	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(conf.Timeout),
	})
	//cors definition
	app.Use(cors.New())

	//routing endpoint
	router.GetRouting(app, db)

	//there is not endpoint
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "endpoint is not found",
			"data":  nil,
		})
	})

	//listen server
	if conf.TmpDep == "local" {
		app.Listen(":" + conf.Port)
	} else {
		app.Listen(":" + config.GetString("PORT"))
	}
}

func connectDatabase(config *config.Config) *gorm.DB {
	db := database.ConnectDB(config)
	return db
}

func migrateTabel(db *gorm.DB) {
	database.AutoMigrate(db, "book", &model.Book{})
	database.AutoMigrate(db, "jenis", &model.Jenis{})
}
