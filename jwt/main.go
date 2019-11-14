package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func public(c *gin.Context) {
	c.JSON(200, gin.H{"message": "This is PUBLIC!!"})
}

func private(c *gin.Context) {
	c.JSON(200, gin.H{"message": "This is PRIVATE!!"})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context)  {
		opt := option.WithCredentialsFile(os.Getenv("CREDENTIAL"))

		app, err := firebase.NewApp(c, nil, opt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Faild to new firebase app (reason: %v)\n", err)
			c.Abort()
		}
		auth, err := app.Auth(c)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Faild to auth client (reason: %v)\n", err)
			c.Abort()
		}
		authHeader := c.GetHeader("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := auth.VerifyIDToken(c, idToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Verify Token is invalid (reason: %v)\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Verify Token is invalid"})
			c.Abort()
		}
		fmt.Println(token.UID)
		fmt.Fprintf(os.Stdout, "Verified Token :%v\n", token)
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"*"},
	}))
	router.GET("/public", public)
	router.GET("/ping", ping)

	authGroup := router.Group("/private")
	authGroup.Use(authMiddleware())
	{
		authGroup.GET("", private)
	}

	router.Run(":8081")
}
