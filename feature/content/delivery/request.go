package delivery

import "portal/domain"

type InsertRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Images  string `json:"images" form:"images"`
	UserID  int    `json:"user_id" form:"user_id"`
}

func (ni *InsertRequest) ToDomain() domain.Content {
	return domain.Content{
		Title:   ni.Title,
		Content: ni.Content,
		UserID:  ni.UserID,
		Images:  ni.Images,
	}
}
