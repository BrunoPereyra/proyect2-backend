package middleware

import (
	"backend/jwt"
	"strings"
    "fmt"
	"github.com/gofiber/fiber/v2"
)

func UseExtractor() fiber.Handler {

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
       
		token := strings.Replace(authHeader, "Bearer ", "", 1)

		nameUser, err := jwt.ExtractDataFromToken(token)
		fmt.Println(nameUser,"name")
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"Message": "Unauthorized",
			})
		}
		c.Context().SetUserValue("nameUser", nameUser)
		return c.Next()
	}

}
