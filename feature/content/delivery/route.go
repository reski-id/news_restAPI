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
	e.POST("/mypost", bc.InsertContent(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/mypost/:id", bc.UpdateContent(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/mypost/:id", bc.DeleteContent(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/allpost", bc.GetAllContent())
	e.GET("/mypost/:id", bc.GetContentID(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/mypost", bc.GetMYPost(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
}
