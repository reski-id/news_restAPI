package data

import (
	"fmt"
	"log"
	"portal/domain"

	"gorm.io/gorm"
)

type contentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ContentData {
	return &contentData{
		db: db,
	}
}

func (nd *contentData) Insert(newData domain.Content) domain.Content {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Content{}
	}
	return cnv.ToDomain()
}

func (bd *contentData) Update(contentID int, updatedData domain.Content) domain.Content {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", contentID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Content{}
	}
	cnv.ID = uint(contentID)
	return cnv.ToDomain()
}

func (nd *contentData) Delete(contentID int) bool {
	err := nd.db.Where("ID = ?", contentID).Delete(&Content{})
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

func (nd *contentData) GetAll() []domain.Content {
	var data []Content
	err := nd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *contentData) GetContentID(contentID int) []domain.Content {
	var data []Content
	err := nd.db.Where("ID = ?", contentID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
