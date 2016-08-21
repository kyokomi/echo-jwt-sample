package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const secretKey = "secret"

func hs256Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func hs256Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func hs256(e *echo.Echo) *echo.Echo {
	e.Post("/hs/login", hs256Login)

	// Restricted group
	r := e.Group("/hs256")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:       func(echo.Context) bool { return false },
		SigningMethod: middleware.AlgorithmHS256,
		ContextKey:    "user",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		SigningKey:    []byte(secretKey),
	}))
	r.GET("/restricted", hs256Restricted)

	return e
}
