package main

import (
	"basic_gin/controller"
	"basic_gin/handle"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	m := handle.NewMember()

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
