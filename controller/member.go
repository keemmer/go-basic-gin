package controller

import (
	"basic_gin/handle"
	"basic_gin/model"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AllData(m *handle.MemberData) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, m.AllData())
	}
}

func AddData(m *handle.MemberData) func(c *gin.Context) {
	return func(c *gin.Context) {
		var v model.Member
		validate := validator.New()
		validate.RegisterValidation("is_keemmer", ValidateKeemmer)

		err := c.ShouldBindJSON(&v)
		if err != nil {
			// c.JSON(400, gin.H{"error": "Error format"})
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = validate.Struct(v)
		if err != nil {
			// c.JSON(400, gin.H{"error": "Error format"})
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		m.AddData(v)
		c.JSON(201, v)
	}
}


func ValidateKeemmer(field validator.FieldLevel) bool {
	fmt.Println("sdfsadfsd")
	return strings.Contains(field.Field().String(), "keemmer")
}
