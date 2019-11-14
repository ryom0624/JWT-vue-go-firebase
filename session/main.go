package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gorilla/securecookie"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	router := gin.Default()
	store := cookie.NewStore(securecookie.GenerateRandomKey(64))
	router.Use(sessions.Sessions("SessionID", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"*"},
	}))

	router.GET("/ping", ping)

	authGroup := router.Group("/session")
	authGroup.Use(authMiddleware())
	{
		authGroup.GET("", session)
	}


	router.Run(":8083")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		fmt.Println("hoge")
		token, err := auth.VerifyIDToken(c, idToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Verify Token is invalid (reason: %v)\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Verify Token is invalid"})
			c.Abort()
		}
		fmt.Println(token.UID)
		c.Set("UID", token.UID)
		fmt.Fprintf(os.Stdout, "Verified Token :%v\n", token)
		c.Next()
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func session(c *gin.Context) {
	uid, _ := c.Get("UID")
	fmt.Println("UID:", uid)
	c.JSON(200, gin.H{"ユーザーID": uid})
}
