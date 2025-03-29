package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/HoneySinghDev/go-echo-rest-api-template/internal/db"
	"github.com/HoneySinghDev/go-echo-rest-api-template/pkg/server"
)

// SignupRequest defines the expected JSON structure for signup.
type SignupRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

// SignupResponse defines the JSON response that includes a JWT token.
type SignupResponse struct {
	Message string      `json:"message"`
	UserID  pgtype.UUID `json:"user_id"`
	Token   string      `json:"token"`
}

func HandleSignupCreate(s *server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(SignupRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		creds := UserCreds{
			EmailID:         req.Email,
			Password:        req.Password,
			ConfirmPassword: req.ConfirmPassword,
		}

		if errs, ok := creds.Validate(); !ok {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errs})
		}

		// Check if the user already exists using the provided email.
		existsResult, err := s.Queries.CheckUserExists(c.Request().Context(), req.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Server error"})
		}
		if existsResult {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "User already exists"})
		}

		// Hash the password.
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
		}

		// Prepare parameters for creating a user.
		createParams := db.CreateUserParams{
			Username:     req.Email, // Using email as username (adjust as needed)
			PasswordHash: string(hashed),
			Email:        req.Email,
		}

		// Create the user in PostgreSQL.
		userID, err := s.Queries.CreateUser(c.Request().Context(), createParams)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
		}

		// Generate JWT token using s.ManagementServerConfig.Secret as the secret key.
		secretKey := []byte(s.Config.Management.Secret)
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = req.Email
		claims["user_id"] = userID
		claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
		}

		return c.JSON(http.StatusOK, SignupResponse{
			Message: "User created successfully",
			UserID:  userID,
			Token:   tokenString,
		})
	}
}
