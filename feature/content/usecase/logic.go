package usecase

import (
	"errors"
	"portal/domain"
)

type contentUseCase struct {
	contentData domain.ContentData
}

func New(model domain.ContentData) domain.ContentUseCase {
	return &contentUseCase{
		contentData: model,
	}
}

func (nu *contentUseCase) AddContent(IDUser int, newContent domain.Content) (domain.Content, error) {
	if IDUser == -1 {
		return domain.Content{}, errors.New("invalid user")
	}

	newContent.UserID = IDUser
	res := nu.contentData.Insert(newContent)

	if res.ID == 0 {
		return domain.Content{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *contentUseCase) UpContent(IDContent int, updateData domain.Content) (domain.Content, error) {
	if IDContent == -1 {
		return domain.Content{}, errors.New("invalid data")
	}

	res := nu.contentData.Update(IDContent, updateData)
	if res.ID == 0 {
		return domain.Content{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *contentUseCase) DelContent(IDContent int) (bool, error) {
	err := nu.contentData.Delete(IDContent)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (nu *contentUseCase) GetAllN() ([]domain.ContentDetail, error) {
	res := nu.contentData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (nu *contentUseCase) GetSpecificContent(newsID int) ([]domain.Content, error) {
	res := nu.contentData.GetContentID(newsID)
	if newsID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}

func (nu *contentUseCase) GetmyContent(IDUser int) ([]domain.ContentDetail, error) {
	res := nu.contentData.GetAllMyContent(IDUser)

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}
