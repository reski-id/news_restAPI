package factory

import (
	ud "portal/feature/user/data"
	userDelivery "portal/feature/user/delivery"
	us "portal/feature/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	nd "portal/feature/content/data"
	contentDelivery "portal/feature/content/delivery"
	nu "portal/feature/content/usecase"

	dd "portal/feature/detail/data"
	detailDelivery "portal/feature/detail/delivery"
	du "portal/feature/detail/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)

	contentData := nd.New(db)
	contentCase := nu.New(contentData)
	contentHandler := contentDelivery.New(contentCase)
	contentDelivery.RouteContent(e, contentHandler)

	detailData := dd.New(db)
	contentDCase := du.New(detailData)
	contentDHandler := detailDelivery.New(contentDCase)
	detailDelivery.RoutePoint(e, contentDHandler)

}
