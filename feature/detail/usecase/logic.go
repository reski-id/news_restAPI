package usecase

import (
	"errors"
	"portal/domain"
)

type contentUseCase struct {
	contentData domain.ContentDetailData
}

func New(model domain.ContentDetailData) domain.ContentDetailUseCase {
	return &contentUseCase{
		contentData: model,
	}
}

func (nu *contentUseCase) AddDetail(newData domain.Detail) (domain.Detail, error) {

	res := nu.contentData.Insert(newData)

	if res.ID == 0 {
		return domain.Detail{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *contentUseCase) UpDetail(IDDetail int, updateData domain.Detail) (domain.Detail, error) {
	if IDDetail == -1 {
		return domain.Detail{}, errors.New("invalid data")
	}

	res := nu.contentData.Update(IDDetail, updateData)
	if res.ID == 0 {
		return domain.Detail{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *contentUseCase) DelDetail(IDDetail int) (bool, error) {
	err := nu.contentData.Delete(IDDetail)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (nu *contentUseCase) GetAllD() ([]domain.Detail, error) {
	res := nu.contentData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (nu *contentUseCase) GetSpecificDetail(newsID int) ([]domain.Detail, error) {
	res := nu.contentData.GetDetailID(newsID)
	if newsID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}
