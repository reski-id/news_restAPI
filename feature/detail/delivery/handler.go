package delivery

import (
	"log"
	"net/http"
	"portal/domain"
	"portal/feature/common"
	"strconv"

	"github.com/labstack/echo/v4"
)

type detailHandler struct {
	contentUsecase domain.ContentDetailUseCase
}

func New(nu domain.ContentDetailUseCase) domain.ContentDetailHandler {
	return &detailHandler{
		contentUsecase: nu,
	}
}

// func (nh *detailHandler) InsertDetail() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var tmp InsertRequest
// 		err := c.Bind(&tmp)

// 		if err != nil {
// 			log.Println("Cannot parse data", err)
// 			c.JSON(http.StatusBadRequest, "error read input")
// 		}

// 		var userid, role = common.ExtractData(c)
// 		if role != "admin" {
// 			return c.JSON(http.StatusCreated, map[string]interface{}{
// 				"message": "Only Admin can Access",
// 			})
// 		}
// 		data, err := nh.contentUsecase.AddDetail(userid, tmp.ToDomain())

// 		if err != nil {
// 			log.Println("Cannot proces data", err)
// 			c.JSON(http.StatusInternalServerError, err)
// 		}

// 		fmt.Println(userid)

// 		return c.JSON(http.StatusCreated, map[string]interface{}{
// 			"message": "success create data",
// 			"data":    FromDomain(data),
// 		})

// 	}
// }

func (nh *detailHandler) InsertDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}
		_, role := common.ExtractData(c)

		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can insert point",
			})
		}

		data, err := nh.contentUsecase.AddDetail(tmp.ToDomain())
		if err != nil {
			log.Println("Cannot proces data,ID sudah Ada", err)

			c.JSON(http.StatusInternalServerError, "Cannot proces data,ID sudah Ada")
			return (err)

		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
		})
	}
}

func (nh *detailHandler) UpdateDetail() echo.HandlerFunc {
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

		var _, role = common.ExtractData(c)
		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can Access",
			})
		}

		data, err := nh.contentUsecase.UpDetail(cnv, tmp.ToDomain())

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

func (nh *detailHandler) DeleteDetail() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := nh.contentUsecase.DelDetail(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete data")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data",
		})
	}
}

func (nh *detailHandler) GetAllDetail() echo.HandlerFunc {
	return func(c echo.Context) error {

		var _, role = common.ExtractData(c)
		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can Access",
			})
		}

		data, err := nh.contentUsecase.GetAllD()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "Cannot get data no data found")

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

func (nh *detailHandler) GetDetailID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var _, role = common.ExtractData(c)
		if role != "admin" {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "Only Admin can Access",
			})
		}

		idDetail := c.Param("id")
		id, _ := strconv.Atoi(idDetail)
		data, err := nh.contentUsecase.GetSpecificDetail(id)

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
