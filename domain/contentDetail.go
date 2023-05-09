package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Detail struct {
	ID          int
	Point       int
	ViewsNumber int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	ContentID   int
}

type ContentDetailHandler interface {
	InsertDetail() echo.HandlerFunc
	UpdateDetail() echo.HandlerFunc
	DeleteDetail() echo.HandlerFunc
	GetAllDetail() echo.HandlerFunc
	GetDetailID() echo.HandlerFunc
}

type ContentDetailUseCase interface {
	// AddDetail(ContentID int, useDetail Detail) (Detail, error)
	AddDetail(newData Detail) (Detail, error)
	UpDetail(IDDetail int, updateData Detail) (Detail, error)
	DelDetail(IDDetail int) (bool, error)
	GetAllD() ([]Detail, error)
	GetSpecificDetail(DetailID int) ([]Detail, error)
}

type ContentDetailData interface {
	Insert(insertDetail Detail) Detail
	Update(IDDetail int, updatedDetail Detail) Detail
	Delete(IDDetail int) error
	GetAll() []Detail
	GetDetailID(DetailID int) []Detail
}
