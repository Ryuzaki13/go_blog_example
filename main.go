package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")
	router.Static("/assets", "resources")

	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/blog", blog)
	router.PUT("/blog", putBlog)

	if Connect() == false {
		return
	}

	_ = router.Run("127.0.0.1:8080")
}

func index(c *gin.Context) {
	c.HTML(200, "index", nil)
}

func about(c *gin.Context) {
	c.HTML(200, "about", nil)
}

func blog(c *gin.Context) {
	c.HTML(200, "blog", nil)
}

type Blog struct {
	Caption string
	Content string
	Image   string
}

func putBlog(c *gin.Context) {
	var b Blog
	e := c.BindJSON(&b)
	if e != nil {
		fmt.Println(e)
		return
	}

	CreateBlog(b.Caption, b.Content)
}
