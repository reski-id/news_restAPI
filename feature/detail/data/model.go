package data

import (
	"portal/domain"

	"gorm.io/gorm"
)

type Detail struct {
	gorm.Model
	Point       int     `json:"point" form:"point"`
	ViewsNumber int     `json:"views_number" form:"views_number"`
	ContentID   int     `gorm:"unique" json:"ContentID" form:"ContentID" validate:"required"`
	Content     Content `gorm:"foreignKey:ContentID; references:ID; constraint:OnDelete:CASCADE"`
}

type Content struct {
	gorm.Model
}

func (b *Detail) ToDomain() domain.Detail {
	return domain.Detail{
		ID:          int(b.ID),
		Point:       b.Point,
		ViewsNumber: b.ViewsNumber,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
		ContentID:   b.ContentID,
	}
}

func ParseToArr(arr []Detail) []domain.Detail {
	var res []domain.Detail

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Detail) Detail {
	var res Detail
	res.ID = uint(data.ID)
	res.ContentID = data.ContentID
	res.Point = data.Point
	res.ViewsNumber = data.ViewsNumber
	return res
}
