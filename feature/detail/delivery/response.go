package delivery

import "portal/domain"

type DataResponse struct {
	ID          int `json:"id"`
	Point       int `json:"point"`
	ViewsNumber int `json:"views_number"`
	ContentID   int `json:"content_id"`
}

func FromDomain(data domain.Detail) DataResponse {
	var res DataResponse
	res.ID = int(data.ID)
	res.Point = data.Point
	res.ViewsNumber = data.ViewsNumber
	res.ContentID = int(data.ContentID)
	return res
}
