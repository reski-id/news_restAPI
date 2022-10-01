package delivery

import (
	"portal/config"
	"portal/domain"
	"portal/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutePoint(e *echo.Echo, bc domain.ContentDetailHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/point", bc.InsertDetail(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/point/:id", bc.UpdateDetail(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/point/:id", bc.DeleteDetail(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/point", bc.GetAllDetail())
	e.GET("/point/:id", bc.GetDetailID())
}
