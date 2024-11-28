package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// GenerateJWT generates a JWT token with the given user ID and role
func GenerateJWT(userID int, role string, expirationHours int) (string, error) {
	var JWT_SECRET = os.Getenv("JWT_SECRET")

	// Default to 72 hours if no expiration time is provided
	if expirationHours == 0 {
		expirationHours = 72
	}

	// Set expiration time
	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour).Unix()

	// Create the JWT claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     expirationTime,
	}

	// Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	signedToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

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
