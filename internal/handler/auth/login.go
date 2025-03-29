package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/HoneySinghDev/go-echo-rest-api-template/pkg/server"
)

// LoginRequest defines the expected JSON structure for login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse defines the JSON response with a JWT token.
type LoginResponse struct {
	Token string `json:"token"`
}

var jwtSecret = []byte("my-secret-key") // In production, load this securely

func HandleLoginCreate(s *server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(LoginRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// Validate input using our existing UserCreds validator.
		creds := UserCreds{
			EmailID:  req.Email,
			Password: req.Password,
		}
		if errs, ok := creds.Validate(); !ok {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errs})
		}

		// Fetch the user from PostgreSQL.
		// (Assuming that GetUserByUsername uses the email as the username.)
		user, err := s.Queries.GetUserByUsername(c.Request().Context(), req.Email)
		if err != nil {
			// Return unauthorized if user is not found.
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		}

		// Compare the provided password with the stored hashed password.
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		}

		// Create a JWT token.
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = req.Email
		claims["user_id"] = user.ID
		claims["exp"] = time.Now().Add(24 * time.Hour).Unix() // token expires in 24 hours

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
		}

		return c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
	}
}
