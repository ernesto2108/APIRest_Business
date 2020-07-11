package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/labstack/echo"
)


func SignUpEndPoint(c *echo.Context) {
}

func LogInEndPoint(c *echo.Context) {
}

func LogInGoogleEndPoint(c *echo.Context) {
}


func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}