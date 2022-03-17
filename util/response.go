package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ResponseHTTP(c *fiber.Ctx, responseCode int, err error, data interface{}) error {
	if err == nil {
		return c.Status(responseCode).JSON(fiber.Map{
			"error": false,
			"data":  data,
		})
	} else {
		return eror(c, responseCode, err)
	}

}

func eror(c *fiber.Ctx, responseCode int, err error) error {
	log := logrus.WithFields(logrus.Fields{
		"Method": c.Method(),
		"Host":   c.Hostname(),
		"Path":   c.Path(),
	}).WithField("ResponseCode", responseCode)

	if responseCode < 500 {
		log.Warn(err.Error())
	} else {
		log.Error(err.Error())
	}

	return c.Status(responseCode).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
		"data":  nil,
	})
}
