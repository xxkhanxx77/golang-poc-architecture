package handler

import (
	"baac-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	service *services.CarService
}

func NewCarHandler(gin *gin.Engine, s *services.CarService) {
	h := CarHandler{service: s}
	gin.GET("/car", h.Car)
}

func (h *CarHandler) Car(c *gin.Context) {
	res, err := h.service.Car()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, res)

}
