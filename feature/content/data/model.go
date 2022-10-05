package data

import (
	"portal/domain"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Images  string `json:"images" form:"images"`
	UserID  int
	User    User `gorm:"foreignKey:UserID; references:ID; constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
}

type ContentDetail struct {
	ID          int
	Title       string
	Content     string
	Images      string
	Point       int
	ViewsNumber int
	ContentID   int
	UserID      int
}

func (b *Content) ToDomain() domain.Content {
	return domain.Content{
		ID:        int(b.ID),
		Title:     b.Title,
		Content:   b.Content,
		Images:    b.Images,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		UserID:    b.UserID,
	}
}

func (d *ContentDetail) ToDomainContentDetails() domain.ContentDetail {
	return domain.ContentDetail{
		ID:          int(d.ID),
		Title:       d.Title,
		Content:     d.Content,
		Point:       d.Point,
		ViewsNumber: d.ViewsNumber,
		UserID:      d.UserID,
		Images:      d.Images,
	}
}

func ParseToArr(arr []Content) []domain.Content {
	var res []domain.Content

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ParseToArrContentDetails(arr []ContentDetail) []domain.ContentDetail {
	var res []domain.ContentDetail

	for _, val := range arr {
		res = append(res, val.ToDomainContentDetails())
	}
	return res
}

func ToLocal(data domain.Content) Content {
	var res Content
	res.ID = uint(data.ID)
	res.UserID = data.UserID
	res.Title = data.Title
	res.Content = data.Content
	res.Images = data.Images
	return res
}
