package delivery

import (
	"log"
	"net/http"
	"portal/config"
	"portal/domain"
	"portal/feature/common"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(e *echo.Echo, us domain.UserUseCase) {
	handler := &userHandler{
		userUsecase: us,
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/register", handler.InsertUser())
	e.POST("/login", handler.LogUser())
	e.GET("/users", handler.GetProfile(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/users", handler.DeleteUser(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/users", handler.UpdateUser(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := uh.userUsecase.AddUser(tmp.ToModel())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
			"token":   common.GenerateToken(data.ID, data.Role),
		})
	}
}

func (uh *userHandler) LogUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		err := c.Bind(&userLogin)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		row, data, e := uh.userUsecase.LoginUser(userLogin.LoginToModel())
		if e != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, "username or password incorrect")
		}
		if row == -1 {
			return c.JSON(http.StatusBadRequest, "username or password incorrect")
		}

		token := common.GenerateToken(int(data.ID), data.Role)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success login",
			"token":   token,
			"role":    data.Role,
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		// id := common.ExtractData(c)
		id, _ := common.ExtractData(c)

		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusFound, map[string]interface{}{
			"message": "data found",
			"data":    data,
		})
	}
}

func (uh *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := common.ExtractData(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		_, errDel := uh.userUsecase.DeleteUser(id)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}
		return c.JSON(http.StatusOK, "success to delete user")
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		var tmp InsertFormat
		result := c.Bind(&tmp)

		idUpdate, _ := common.ExtractData(c)

		if result != nil {
			log.Println(result, "Cannot parse input to object")
			return c.JSON(http.StatusInternalServerError, "Error read update")
		}

		data, err := uh.userUsecase.UpdateUser(idUpdate, tmp.ToModel())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Success update data",
			"data":    data,
		})
	}
}
