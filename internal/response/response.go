package response

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, data interface{}) error {
	res := fiber.Map{"code:": 200, "success": true}

	if data != nil {
		res["data"] = data
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func Error(c *fiber.Ctx, statusCode int, message interface{}) error {
	var msg string

	switch v := message.(type) {
	case error:
		msg = v.Error()
	case string:
		msg = v
	}

	res := fiber.Map{"code": statusCode, "success": false}

	if message != "" {
		res["message"] = msg
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
