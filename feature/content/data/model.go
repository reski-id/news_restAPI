package data

import (
	"portal/domain"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Content string `json:"content" form:"content"`
	Images  string `json:"images" form:"images"`
	UserID  int
	User    User `gorm:"foreignKey:UserID; references:ID; constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
}

func (b *Content) ToDomain() domain.Content {
	return domain.Content{
		ID:        int(b.ID),
		Content:   b.Content,
		Images:    b.Images,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		UserID:    b.UserID,
	}
}

func ParseToArr(arr []Content) []domain.Content {
	var res []domain.Content

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Content) Content {
	var res Content
	res.ID = uint(data.ID)
	res.UserID = data.UserID
	res.Content = data.Content
	res.Images = data.Images
	return res
}
