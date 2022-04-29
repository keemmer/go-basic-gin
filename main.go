package main

import (
	"basic_gin/controller"
	"basic_gin/handle"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	m := handle.NewMember()
	// r.Use(Auth)

	// validate := validator.New()
	// validate.RegisterValidation("is_keemmer", ValidateKeemmer)

	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte("<html><body>MY REEST API</body></html>"))

	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/member", controller.AllData(m))
	r.POST("/member", controller.AddData(m))

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	auth := strings.Split(token, " ")
	if len(auth) != 2 || auth[0] != "Bearer" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Authorization failed"})
		return
	}
	if auth[1] != "344wrtso245fklpo408!343123" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization"})
		return
	}
	c.Next()
}

// func ValidateKeemmer(field validator.FieldLevel) bool {
// 	fmt.Println("sdfsadfsd")
// 	return strings.Contains(field.Field().String(), "keemmer")
// }
