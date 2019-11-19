package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type sendData struct {
	Product string `json:product`
}

type User struct {
	UserId string
	Password string
}


func main() {
	router := gin.Default()
	store := cookie.NewStore(securecookie.GenerateRandomKey(64))
	router.Use(sessions.Sessions("SessionID", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/ping", ping)

	authGroup := router.Group("/session")
	authGroup.Use(authMiddleware(), setUIDSession())
	{
		authGroup.GET("/getuid", session)
		authGroup.POST("/cart", addCart)
		authGroup.GET("/items", getItems)
	}


	router.Run(":8089")
}

func authMiddleware() gin.HandlerFunc  {
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

		token, err := auth.VerifyIDToken(c, idToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Verify Token is invalid (reason: %v)\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Verify Token is invalid"})
			c.Abort()
		}

		c.Set("UserID", token.UID)

		fmt.Fprintf(os.Stdout, "Verified Token :%v\n", token)
		c.Next()
	}
}

func setUIDSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		s := session.Get("UserID")
		id, _ := c.Get("UserID")
		if id != s {
			session.Set("UserID", id)
			session.Save()
		}
	}
}
func ping(c *gin.Context) {
	log.Println(c.Request.Header["Cookie"])
	c.JSON(200, gin.H{"message": "pong"})
}

func session(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("UserID")
	c.JSON(200, gin.H{"ユーザーID": uid})
}

func addCart(c *gin.Context) {
	postData := sendData{}
	c.BindJSON(&postData)

	session := sessions.Default(c)

	switch sessionData := session.Get("Products").(type) {
	case nil:
		session.Set("Products", postData.Product)
	case string:
		tempSlice := []string{postData.Product}
		tempSlice = append(tempSlice, sessionData)
		session.Set("Products", tempSlice)
	case []string:
		sessionData = session.Get("Products").([]string)
		sessionData = append(sessionData, postData.Product)
		session.Set("Products", sessionData)
	}

	session.Save()

	//var sessionData []string
	//if session.Get("Products") != nil {
	//	sessionData = session.Get("Products").([]string)
	//	sessionData = append(sessionData, postData.Product)
	//	session.Set("Products", sessionData)
	//} else {
	//	session.Set("Products", postData.Product)
	//}

	//data1 := []string{"hoge", "fuga"}
	//data1 = append(data1, "piyo")
	//session.Set("data1", data1)
	//
	//sdata1 := session.Get("data1")
	////sdata1 = append(sdata1, "foo")
	//
	//fmt.Println(reflect.TypeOf(data1)) // []interface{}
	//fmt.Println(data1)
	//fmt.Println(reflect.TypeOf(sdata1)) // []interface{}
	//fmt.Println(sdata1)

	//var data2 interface{}
	//data2 = "hoge"
	//session.Set("data2", data2)
	//sdata2 := session.Get("data2") // hoge"
	//fmt.Println(reflect.TypeOf(data2))
	//spew.Dump(data2)
	//fmt.Println(reflect.TypeOf(sdata2))
	//spew.Dump(sdata2)

	//for _, x := range data1 {
	//	fmt.Println(reflect.TypeOf(x))
	//}
	//for _, y := range sdata1 {
	//	fmt.Println(reflect.TypeOf(y))
	//}

	c.JSON(200, gin.H{
		"You added to cart": postData.Product,
		//"UserID": session.Get("UserID"),
		//"Products": session.Get("Products"),
		//"data1": session.Get("data1"), // [ "hoge", "fuga", "piyo" ]
	})
}

func getItems(c *gin.Context) {
	session := sessions.Default(c)
	products := session.Get("Products")

	if products == nil {
		c.JSON(200, gin.H{
			"Products": "There is no products.",
		})
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(200, gin.H{
			"Products": products,
		})
	}
}