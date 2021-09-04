package main

import (
	"go-rest-api/app"
	"go-rest-api/db"
	"go-rest-api/domain/posts"
)

func main() {
	db := db.InitDb()
	defer db.Close()

	postRepo := posts.NewPostRepository(db)
	postService := posts.NewPostService(postRepo)
	postController := posts.NewPostController(postService)

	server := app.NewHttpServer()

	posts.NewPostRouter(server, postController)

	server.Logger.Fatal(server.Start((":8000")))
}
