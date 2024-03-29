package data

import (
	"fmt"
	"log"
	"portal/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) Insert(newUser domain.User) (domain.User, error) {
	var cnv = ToLocal(newUser)
	err := ud.db.Create(&cnv).Error
	if err != nil {
		log.Println("Cannot create object", err.Error())
		return domain.User{}, err
	}

	return cnv.ToDomain(), nil
}
func (ud *userData) Update(userID int, updatedData domain.User) domain.User {
	var cnv = ToLocal(updatedData)
	err := ud.db.Model(&User{}).Where("ID = ?", userID).Updates(cnv).Error
	if err != nil {
		log.Println("Cannot update data", err.Error())
		return domain.User{}
	}
	cnv.ID = uint(userID)
	return cnv.ToDomain()
}

func (ud *userData) Login(userLogin domain.User) (row int, data domain.User, err error) {
	var dataUser = ToLocal(userLogin)
	password := dataUser.Password

	result := ud.db.Where("username = ?", dataUser.Username).First(&dataUser)

	if result.Error != nil {
		return 0, domain.User{}, result.Error
	}

	if result.RowsAffected != 1 {
		return -1, domain.User{}, fmt.Errorf("failed to login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(password))

	if err != nil {
		return -2, domain.User{}, fmt.Errorf("failed to login")
	}

	return int(result.RowsAffected), dataUser.ToDomain(), nil
}

func (ud *userData) GetSpecific(userID int) (domain.User, error) {
	var tmp User
	err := ud.db.Where("ID = ?", userID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.User{}, err
	}

	return tmp.ToDomain(), nil
}

func (ud *userData) Delete(userID int) (row int, err error) {
	res := ud.db.Delete(&User{}, userID)
	if res.Error != nil {
		log.Println("Cannot delete data", res.Error.Error())
		return 0, res.Error
	}

	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return 0, fmt.Errorf("failed to delete user")
	}
	return int(res.RowsAffected), nil
}
func (ud *userData) GetAll() []domain.User {
	var data []User
	// err := ud.db.Find(&data)
	err := ud.db.Where("role != ?", "admin").Find(&data)
	// err := ud.db.Select("full_name", "email").Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr2(data)
}
