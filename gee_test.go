package gee

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestServer(t *testing.T) {
	Server()
}

func TestGf(t *testing.T) {
	g.Server().Run()
}

func TestGin(t *testing.T)  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000")
}

func TestGee(t *testing.T) {
	r := New()
	t.Log(r)
}