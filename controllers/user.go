package controllers

import (
	"net/http"

	"github.com/denisemignoli/to-do-list/models"
	"github.com/denisemignoli/to-do-list/repositories"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepository repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{
		UserRepository: repo,
	}
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uc.UserRepository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user.ID = id
	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserRepository.GetUserByUsername(input.Username)
	if err != nil || user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
