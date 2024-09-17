package validation

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func Validate[T interface{}](c *gin.Context) (*T, error) {
	var input T

	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	v := validator.New()

	err := v.Struct(input)

	if err != nil {
		slog.Error("Error binding body", err)
		return nil, err
	}

	return &input, nil
}
