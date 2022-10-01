package data

import (
	"fmt"
	"log"
	"portal/domain"

	"gorm.io/gorm"
)

type contentDetailData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ContentDetailData {
	return &contentDetailData{
		db: db,
	}
}

func (nd *contentDetailData) Insert(newData domain.Detail) domain.Detail {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Detail{}
	}
	return cnv.ToDomain()
}

func (bd *contentDetailData) Update(dataID int, updatedData domain.Detail) domain.Detail {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", dataID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Detail{}
	}
	cnv.ID = uint(dataID)
	return cnv.ToDomain()
}

func (nd *contentDetailData) Delete(dataID int) bool {
	err := nd.db.Where("ID = ?", dataID).Delete(&Content{})
	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}

func (nd *contentDetailData) GetAll() []domain.Detail {
	var data []Detail
	err := nd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *contentDetailData) GetDetailID(dataID int) []domain.Detail {
	var data []Detail
	err := nd.db.Where("ID = ?", dataID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
