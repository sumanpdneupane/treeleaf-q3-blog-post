package middleware

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware checks if the user is authenticated via JWT
func AuthMiddleware(c *fiber.Ctx) error {
	// Secret key used for signing JWT tokens. In a real application, store this securely.
	var JWT_SECRET = os.Getenv("JWT_SECRET")

	token := c.Get("Authorization")
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}

	// Remove the "Bearer " prefix if present
	token = strings.Replace(token, "Bearer ", "", 1)

	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil // Using the shared secret for HMAC
	})

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Error parsing token: "+err.Error())
	}

	if !parsedToken.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	c.Locals("user", claims)
	return c.Next()
}
