package delivery

import "portal/domain"

type InsertFormat struct {
	UserName string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	FullName string `json:"fullname" form:"fullname"`
	Role     string `json:"role" form:"role"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		UserName: i.UserName,
		Email:    i.Email,
		Password: i.Password,
		FullName: i.FullName,
		Role:     i.Role,
	}
}

type LoginFormat struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (lf *LoginFormat) LoginToModel() domain.User {
	return domain.User{
		UserName: lf.UserName,
		Password: lf.Password,
	}
}
