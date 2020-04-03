package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kazune-br/gin-gorm-firebase-minimum-template/internal/configuration"
)

func main() {
	configuration.Load()
	ctx := context.Background()

	// initialize firebase app
	credentials, err := google.CredentialsFromJSON(ctx, []byte(configuration.Get().FireBaseJson))
	if err != nil {
		log.Println("failed to read firebase service key")
		panic(err)
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Println("failed to initialize firebase app")
		panic(err)
	}

	// generate firebase auth
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Println("failed to create firebase auth client")
		panic(err)
	}

	// validate token
	token := ""
	jwt, err := auth.VerifyIDToken(ctx, token)
	if err != nil {
		log.Printf("invalid token is given \nerr: %s", err)
	}

	if jwt != nil {
		log.Println("succeed in reading uid")
		log.Println(jwt.UID)
	}

	// initialize db
	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		configuration.Get().DBUser,
		configuration.Get().DBPass,
		configuration.Get().DBHost,
		configuration.Get().DBPort,
		configuration.Get().DBName,
	)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Printf("fail to initialize db")
		panic(err)
	}
	db.LogMode(true)
	defer db.Close()

	// Sample CRUD Commands for Gorm
	// Create
	// u := User{UID: uint64(1)}
	// db.Create(&u)

	// Read
	// users := []User{}
	// db.Find(&users)

	// Delete
	// db.Where("UID = ?", uint64(1)).Delete(&User{})

	// initialize api server
	router := gin.Default()
	router.Use(cors.Default())

	// register an endpoint for health check
	router.GET("/api/healthcheck", HealthCheck)

	// run server
	if err = router.Run(fmt.Sprintf(":%d", configuration.Get().AppPort)); err != nil {
		log.Printf("failed to initialize server")
		panic(err)
	}
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "health check passed"})
}

type User struct {
	// gorm.Model
	// https://gorm.io/ja_JP/docs/conventions.html
	ID  uint64 `gorm:"primary_key"`
	UID uint64 `json:"uid"`
}
