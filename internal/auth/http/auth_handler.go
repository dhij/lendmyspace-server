package http

import "github.com/gin-gonic/gin"

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}
