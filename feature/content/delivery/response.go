package delivery

import "portal/domain"

type DataResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Images  string `json:"images"`
	UserID  int    `json:"user_id"`
}

func FromDomain(data domain.Content) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.Content = data.Content
	res.Images = data.Images
	res.UserID = int(data.UserID)
	return res
}