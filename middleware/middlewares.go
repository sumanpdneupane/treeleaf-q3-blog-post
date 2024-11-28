package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Secret key used for signing JWT tokens. In a real application, store this securely.
var JWT_SECRET = os.Getenv("JWT_SECRET")

// AuthMiddleware checks if the user is authenticated via JWT
func AuthMiddleware(c *fiber.Ctx) error {
	// Extract the token from the Authorization header
	token := c.Get("Authorization")
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}

	// Remove the "Bearer " prefix if present
	token = strings.Replace(token, "Bearer ", "", 1)

	// Parse and validate the token
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key for validation
		return []byte(JWT_SECRET), nil
	})

	fmt.Println(parsedToken.Valid)
	if err != nil || !parsedToken.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token"+err.Error())
	}

	// Optionally, you can add the claims to the request context
	c.Locals("user", claims)

	// Allow the request to proceed to the next handler
	return c.Next()

}
