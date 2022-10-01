package delivery

import "portal/domain"

type InsertRequest struct {
	Content string `json:"content" form:"content"`
	Images  string `json:"images" form:"images"`
	UserID  int    `json:"user_id" form:"user_id"`
}

func (ni *InsertRequest) ToDomain() domain.Content {
	return domain.Content{
		Content: ni.Content,
		UserID:  ni.UserID,
		Images:  ni.Images,
	}
}
