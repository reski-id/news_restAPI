package delivery

import "portal/domain"

type InsertRequest struct {
	Point       int `json:"point" form:"point"`
	ViewsNumber int `json:"views_number" form:"views_number"`
	ContentID   int `json:"content_id" form:"content_id"`
}

func (ni *InsertRequest) ToDomain() domain.Detail {
	return domain.Detail{
		Point:       ni.Point,
		ViewsNumber: ni.ViewsNumber,
		ContentID:   ni.ContentID,
	}
}
