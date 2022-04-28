package main

import (
	"basic_gin/handle"
	"basic_gin/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	m := handle.NewMember()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/member", func(c *gin.Context) {
		c.JSON(200, m.AllData())
	})
	r.POST("/member", func(c *gin.Context) {
		var v model.Member
		err := c.ShouldBindJSON(&v)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error format",
			})
			return
		}
		m.AddData(v)
		c.JSON(201, v)
	})

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
