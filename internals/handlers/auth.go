package handlers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/shubGupta10/shared-space-server/internals/config"
	"github.com/shubGupta10/shared-space-server/internals/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	//parse the rquest body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	//validate the request body
	if data["email"] == "" || data["password"] == "" || data["name"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide all required fields: name, email, and password."})
	}

	//check if user exist or not in db
	var existingUser models.User
	if err := config.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	}

	//hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	//create a new user
	user := models.User{
		ID:        uuid.New(),
		Name:      data["name"],
		Email:     data["email"],
		Password:  string(password),
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(user)

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	//parse the request body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	//validate the request body
	email := data["email"]
	password := data["password"]

	if email == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide both email and password."})
	}

	//check if user exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", email).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User does not exist"})
	}

	//check if password correct or not
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Password is incorrect"})
	}

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  existingUser.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	//get secret from env for jwt
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}

	return c.JSON(fiber.Map{"token": tokenString, "user": fiber.Map{
		"id":    existingUser.ID,
		"name":  existingUser.Name,
		"email": existingUser.Email,
	}})
}

func GetProfile(c *fiber.Ctx) error {

	//take id from auth middleware
	userId := c.Locals("user_id")
	if userId == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Fetch user from DB
	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// don't return password
	user.Password = ""

	return c.JSON(user)
}
