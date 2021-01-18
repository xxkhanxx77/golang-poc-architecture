package handler

import (
	"baac-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(gin *gin.Engine, s *services.UserService) {
	h := UserHandler{service: s}
	gin.GET("/user", h.User)
}

func (h *UserHandler) User(c *gin.Context) {
	res, err := h.service.User()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, res)

}
