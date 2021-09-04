package main

import (
	"fmt"
	"go-rest-api/app"
	"go-rest-api/config"
	"go-rest-api/db"
	"go-rest-api/domain/posts"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Couldn't load config ", err)
	}

	for _, o := range config.Host.AllowOrigins {
		fmt.Println(o)
	}

	fmt.Println()

	db := db.InitDb(config)
	defer db.Close()

	postRepo := posts.NewPostRepository(db)
	postService := posts.NewPostService(postRepo)
	postController := posts.NewPostController(postService)

	server := app.New(config)

	posts.NewPostRouter(server, postController)

	server.Logger.Fatal(server.Start((":" + config.Host.Port)))
}
