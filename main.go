package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"privex/database"
	"privex/handler"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var db = make(map[string]string)

func setupRouter(queries *database.Queries) *gin.Engine {

	// create database struct
	env := handler.Env{Db: queries}

	// create logger
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(gin.Logger())

	// Static files
	// r.Static("/", "./assets")
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("./assets"))))

	api := r.Group("/api")

	// Ping test
	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api.GET("/message", env.ListMessages)
	api.GET("/message/:id", env.GetMessage)
	api.POST("/message", env.PostMessage)
	api.DELETE("/message/:id", env.DeleteMessage)

	api.GET("/user/:id", env.GetUser)
	api.POST("/user", env.PostUser)

	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}

func main() {
	// setup DB
	conn_string := os.Getenv("DB_CONN")
	if conn_string == "" {
		panic("No db connection string")
	}

	conn, err := pgx.Connect(context.Background(), conn_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	queries := database.New(conn)

	// queries.CreateMessage(context.Background(), database.CreateMessageParams{UserID: 1, MessageText: "Some messate", MessageType: "txt"})
	// messages, err := queries.ListMessages(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(messages)
	// }

	// run listener
	r := setupRouter(queries)
	r.Run(":8080")
}
