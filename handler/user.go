package handler

import (
	"baac-backend/entity"
	"baac-backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	service *services.UserService
}

type userRequest struct {
	Name        string `json:"name" `
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}
type UserResponse struct {
	Name        string `json:"Name"`
	PhoneNumber string `json:"PhoneNumber"`
	Email       string `json:"email"`
}

func NewUserHandler(gin *gin.Engine, s *services.UserService) {
	h := UserHandler{service: s}

	gin.GET("/user", h.User)
	gin.POST("/user", h.CreateUser)

}

func (h *UserHandler) User(c *gin.Context) {
	res, err := h.service.User()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, res)

}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req userRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// user := req.MapToUserEntity()
	ent := entity.User{
		ID:          primitive.NewObjectID(),
		Name:        &req.Name,
		PhoneNumber: &req.PhoneNumber,
		Email:       &req.Email,
	}
	fmt.Println(ent)
	res, err := h.service.CreateUser(&ent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return

	}
	result := MapToUserRes(res)

	c.JSON(http.StatusCreated, result)

	return
}

func MapToUserRes(v *entity.User) UserResponse {

	return UserResponse{
		Name:        *v.Name,
		PhoneNumber: *v.PhoneNumber,
		Email:       *v.Email,
	}
}
