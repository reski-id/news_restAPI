package delivery

import "portal/domain"

type InsertRequest struct {
	Content     string `json:"content" form:"content"`
	ViewsNumber int    `json:"view_number" form:"view_number"`
	Point       int    `json:"point" form:"point"`
	Images      string `json:"images" form:"images"`
	UserID      int    `json:"user_id" form:"user_id"`
}

func (ni *InsertRequest) ToDomain() domain.Content {
	return domain.Content{
		Content:     ni.Content,
		ViewsNumber: ni.ViewsNumber,
		Point:       ni.Point,
		UserID:      ni.UserID,
		Images:      ni.Images,
	}
}
