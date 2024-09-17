package controllers

import (
	"fmt"
	"log/slog"
	"net/http"
	"server/src/models"
	"server/src/utils"
	"server/src/validation"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var jwt = new(utils.JWT)

// Generic validate function

func (AuthController) SignUp(c *gin.Context) {

	input, err := validation.Validate[validation.SignupForm](c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	password, err := utils.Password(input.Password).HashPassword()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
	}

	u := models.User{
		Email:    input.Email,
		Password: password,
		Name:     input.Name,
	}

	u.Create()

	c.JSON(http.StatusOK, gin.H{
		"id": u.ID,
	})
}

func (AuthController) Login(c *gin.Context) {

	input, err := validation.Validate[validation.LoginForm](c)

	fmt.Println("input: ", input, err)

	user := models.User{}

	err = user.GetByEmail(input.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Email or Password",
		})
		return
	}

	fmt.Println("Record Found", user)

	IsPasswordMatched := utils.Password(input.Password).ComparePassword(user.Password)

	if !IsPasswordMatched {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Email or Password",
		})
		return
	}

	slog.Info("Generating User by email")

	token, err := jwt.GenerateToken(user.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Token failed",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
