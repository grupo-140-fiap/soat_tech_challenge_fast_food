package handlers

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}
