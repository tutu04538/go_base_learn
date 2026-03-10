package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义一个 middleware
// 该 middleware 会在请求到达时执行
// ctx.Next() 会执行业务代码
// 执行完后返回 middleware
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		latency := time.Since(t)
		log.Printf("time: %v", latency)
	}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "tutu1", Age: 13},
	{ID: 2, Name: "tutu2", Age: 23},
}

func main() {
	r := gin.Default()
	r.Use(Logger())

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/user/:name", func(ctx *gin.Context) {
	// 	name := ctx.Param("name")
	// 	ctx.String(http.StatusOK, "hello %s", name)
	// })

	// r.GET("/users", func(ctx *gin.Context) {
	// 	name := ctx.Query("name")
	// 	age := ctx.DefaultQuery("age", "20")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name": name,
	// 		"age":  age,
	// 	})
	// })

	// r.POST("/form", func(ctx *gin.Context) {
	// 	username := ctx.PostForm("username")
	// 	password := ctx.DefaultPostForm("password", "")

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"username": username,
	// 		"password": password,
	// 	})
	// })

	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserByID)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for _, user := range users {
		if id == user.ID {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if id == user.ID {
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	for i, user := range users {
		if id == user.ID {
			// ... 是可变参数操作符
			// 在调用函数时放在切片后面，表示将切片拆解为独立的参数
			// 在定义函数时，放在类型前，表示允许函数接收任意数量的参数
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
