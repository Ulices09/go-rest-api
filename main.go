package main

import (
	"go-rest-api/internal/app"
	"go-rest-api/internal/config"
	"go-rest-api/internal/db"
	"go-rest-api/internal/domain/auth"
	"go-rest-api/internal/domain/posts"
	"go-rest-api/internal/domain/users"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Couldn't load config ", err)
	}

	db := db.InitDb(config)
	defer db.Close()

	cMiddleware := app.InitMiddlware(config)
	server := app.New(config, cMiddleware)

	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo, config)
	authController := auth.NewAuthController(authService)
	auth.NewAuthRouter(server, authController, cMiddleware)

	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)
	userController := users.NewUserController(userService)
	users.NewUserRouter(server, userController)

	postRepo := posts.NewPostRepository(db)
	postService := posts.NewPostService(postRepo)
	postController := posts.NewPostController(postService)
	posts.NewPostRouter(server, postController, cMiddleware)

	server.Logger.Fatal(server.Start((":" + config.Host.Port)))
}
