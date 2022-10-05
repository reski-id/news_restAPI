package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Content struct {
	ID        int
	Title     string
	Content   string
	Images    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    int
}

type ContentDetail struct {
	ID          int
	Title       string
	Content     string
	Images      string
	Point       int
	ViewsNumber int
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type ContentHandler interface {
	InsertContent() echo.HandlerFunc
	UpdateContent() echo.HandlerFunc
	DeleteContent() echo.HandlerFunc
	GetAllContent() echo.HandlerFunc
	GetMYPost() echo.HandlerFunc
	GetContentID() echo.HandlerFunc
}

type ContentUseCase interface {
	AddContent(IDUser int, useContent Content) (Content, error)
	UpContent(IDContent int, updateData Content) (Content, error)
	DelContent(IDContent int) (bool, error)
	GetAllN() ([]ContentDetail, error)
	GetSpecificContent(ContentID int) ([]Content, error)
	GetmyContent(userID int) ([]ContentDetail, error)
}

type ContentData interface {
	Insert(insertContent Content) Content
	Update(IDContent int, updatedContent Content) Content
	Delete(IDContent int) bool
	GetAll() []ContentDetail
	GetAllMyContent(IDUser int) []ContentDetail
	GetContentID(ContentID int) []Content
}
