package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luccasniccolas/monitor/data"
	"github.com/luccasniccolas/monitor/repositories"
)

func SignUp(c *fiber.Ctx) error {
	user := new(data.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"mesagge": "error al parsear datos",
		})
	}

	err := repositories.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "usuario creado exitosamente",
	})

}

func SignIn(c *fiber.Ctx) error {
	user := new(data.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"mesagge": "error al parsear datos",
		})
	}
	loginUser, err := repositories.Login(user.Email, user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sesión iniciada correctamente",
		"data":    loginUser,
	})
}
