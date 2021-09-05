package main

import (
	"go-rest-api/app"
	"go-rest-api/config"
	"go-rest-api/db"
	"go-rest-api/domain/auth"
	"go-rest-api/domain/posts"
	"go-rest-api/domain/users"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Couldn't load config ", err)
	}

	db := db.InitDb(config)
	defer db.Close()

	server := app.New(config)

	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authController := auth.NewAuthController(authService)
	auth.NewAuthRouter(server, authController)

	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)
	userController := users.NewUserController(userService)
	users.NewUserRouter(server, userController)

	postRepo := posts.NewPostRepository(db)
	postService := posts.NewPostService(postRepo)
	postController := posts.NewPostController(postService)
	posts.NewPostRouter(server, postController)

	server.Logger.Fatal(server.Start((":" + config.Host.Port)))
}
