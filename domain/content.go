package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Content struct {
	ID        int
	Content   string
	Images    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    int
}

type ContentHandler interface {
	InsertContent() echo.HandlerFunc
	UpdateContent() echo.HandlerFunc
	DeleteContent() echo.HandlerFunc
	GetAllContent() echo.HandlerFunc
	GetContentID() echo.HandlerFunc
}

type ContentUseCase interface {
	AddContent(IDUser int, useContent Content) (Content, error)
	UpContent(IDContent int, updateData Content) (Content, error)
	DelContent(IDContent int) (bool, error)
	GetAllN() ([]Content, error)
	GetSpecificContent(ContentID int) ([]Content, error)
}

type ContentData interface {
	Insert(insertContent Content) Content
	Update(IDContent int, updatedContent Content) Content
	Delete(IDContent int) bool
	GetAll() []Content
	GetContentID(ContentID int) []Content
}
