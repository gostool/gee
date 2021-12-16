package gee

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"net/http"
	"testing"
)

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
	r.GET("/", func(c *Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *Context) {
		// except /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *Context) {
		c.JSON(http.StatusOK, H{
			"username":c.PostForm("username"),
			"password":c.PostForm("password"),
		})
	})
	r.Run(":8000")
}