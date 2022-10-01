package delivery

import (
	"portal/config"
	"portal/domain"
	"portal/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteContent(e *echo.Echo, bc domain.ContentHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/post", bc.InsertContent(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/post/:id", bc.UpdateContent(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/post/:id", bc.DeleteContent(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/post", bc.GetAllContent())
	e.GET("/post/:id", bc.GetContentID())
}
