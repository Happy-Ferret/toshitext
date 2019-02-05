package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func SetCookieMiddlewares(g *echo.Group) {
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "You don't have any cookie foo!")
			}
			log.Println(err)
			return err
		}

		if cookie.Value == "random_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "You don't have the right chocolate cookie")

	}
}