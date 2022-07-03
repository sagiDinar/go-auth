package controllers

import (
	"bytes"
	"fibr/controllers/models"
	"fibr/models"
	"go/token"

	"strconv"
	"time"

	"../database"
	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), cost:14)
	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,

	}

	database.DB.Create(&user)
	return c.JSON(user)

	func Login(c *fiber.Ctx) error {
		var data map[string]string
		if err := c.BodyParser(&data); err != nil {
			return err
		}

		var user models.User

		database.DB.Where(query:"email = ?", data["email"]).First(&user)

		if user.Id == 0 {
			c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"massage": "Incorrect username or password", 
		})
		}

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"massage": "Incorrect username or password"
			})
		}

		claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
			Issuer: srtconv.Itoa(int(user.Id)),
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		})

		token, err := claims.SignedString([]byte(SecretKey))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"massage": "Could not login",
			})
		}

		cookie := fiber.Cookie{
			Name: "jwt",
			Value: token,
			Expires:time.Now().Add(time.Hour * 1),
			HTTPOnly: true,
			SameSite: Lax,
		}

		c.Cookie(&cookie)

		retun c.JSON(fiber.Map{
			"massage":"success",
		})
		
	}

	func User(c *fiber.Ctx) error {
		cookie = c.Cookies(key: "jwt")

		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil {
			c.Status(fiber.Map{
				"massage":"unauthenticated",
			})
			
		}

		claims := token.claims.(*jwt.StandardClaims)

		var user models.User

		database.DB.Where(query:"id = ?", claims.Issuer).First(&user)

		return c.JSON(user)
	}
}
