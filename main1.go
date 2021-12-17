package main

import (
	"github.com/gin-gonic/gin"
)

// func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/welcome", welcome)
	r.GET("/user/:name", para)
	r.GET("/homepage", home_page)
	r.POST("/post", post_test)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func para(c *gin.Context) {
	name := c.Param("name")
	c.String(200, name)
}
func welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Gary")
	lastname := c.Query("lastname")
	c.String(200, "hello %v %v", firstname, lastname)
}

// curl -X POST --data "nick=gary" http://0.0.0.0:8080
func post_test(c *gin.Context) {
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{
		"status": "Posted",
		"nick":   nick,
	})
}
func home_page(c *gin.Context) {
	c.HTML(200, "example.html", nil)
}
