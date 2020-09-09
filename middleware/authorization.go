package middleware

import (
	"net/http"

	"github.com/Leonardo-Antonio/golang-echo/certificates/authorization"
	"github.com/labstack/echo"
)

// Authorization .
func Authorization(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			forbidden := map[string]string{"error": "No autorizado"}
			return c.JSON(http.StatusForbidden, forbidden)
		}
		return f(c)
	}
}
