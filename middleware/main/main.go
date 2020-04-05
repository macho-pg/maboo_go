package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Name string		`form:"name"`
	Age  int        `form:"age"`
}

func middlewareCheck() gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println("do login check")
		c.Next()
		fmt.Println("login check end")
	}
}

func index(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "index",
	})
}

func data(c *gin.Context){
	var u user
	fmt.Println("post data")
	if err := c.ShouldBind(&u); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"name": u.Name,
			"age": u.Age,
		})
	}
}

func main() {
	r := gin.Default()
	userGroup := r.Group("/user", middlewareCheck())
	{
		userGroup.GET("/index", index)
		userGroup.POST("/data", data)
	}
	r.Run(":9000")
}