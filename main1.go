package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/welcome", welcome)
	r.GET("/user/:name", para)
	r.POST("/post", post_test)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"haha": "肥豬",
	})
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
func post_test(c *gin.Context) {
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{
		"status": "Posted",
		"nick":   nick,
	})
}
