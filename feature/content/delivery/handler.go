package delivery

import (
	"fmt"
	"log"
	"net/http"
	"portal/domain"
	"portal/feature/common"
	"strconv"

	"github.com/labstack/echo/v4"
)

type contentHandler struct {
	contentUsecase domain.ContentUseCase
}

func New(nu domain.ContentUseCase) domain.ContentHandler {
	return &contentHandler{
		contentUsecase: nu,
	}
}

func (nh *contentHandler) InsertContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		fmt.Println(tmp)

		var userid, _ = common.ExtractData(c)
		data, err := nh.contentUsecase.AddContent(userid, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Println(userid)

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})

	}
}

func (nh *contentHandler) UpdateContent() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp InsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}
		data, err := nh.contentUsecase.UpContent(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"data":    data,
		})
	}
}

func (nh *contentHandler) DeleteContent() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		_, errResult := nh.contentUsecase.DelContent(cnv)
		if errResult != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  errResult.Error(),
				"message": "Cannot deleted data",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data",
		})
	}
}

func (nh *contentHandler) GetAllContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := nh.contentUsecase.GetAllN()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all data",
			"data":    data,
		})
	}
}

func (nh *contentHandler) GetMYPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		// idNews := c.Param("id")
		userid, _ := common.ExtractData(c)
		fmt.Println(userid)
		data, err := nh.contentUsecase.GetmyContent(userid)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusNotFound, "data empty, anda belum ada insert news")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my content",
			"data":    data,
		})
	}
}

func (nh *contentHandler) GetContentID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idContent := c.Param("id")
		id, _ := strconv.Atoi(idContent)
		data, err := nh.contentUsecase.GetSpecificContent(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my data",
			"users":   data,
		})
	}
}
