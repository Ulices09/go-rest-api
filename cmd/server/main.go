package main

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/infrastructure/db"
	"go-rest-api/internal/infrastructure/logger"
	httpapp "go-rest-api/internal/interface/http"
	"go-rest-api/internal/modules/auth"
	"go-rest-api/internal/modules/posts"
	"go-rest-api/internal/modules/users"
	"log"
)

func main() {
	config, err := config.Load(".")

	if err != nil {
		log.Fatal("Couldn't load config ", err)
		return
	}

	logger, err := logger.New(config)

	if err != nil {
		log.Fatal("Couldn't initialize logger ", err)
		return
	}

	defer logger.Sync()

	db, err := db.New(config)

	if err != nil {
		logger.Fatal("Couldn't initialize database ", err)
		return
	}

	defer db.Close()

	cMiddleware := httpapp.InitMiddlware(config)
	server := httpapp.New(config, cMiddleware)

	authRepo := auth.NewAuthRepository(db, logger)
	authService := auth.NewAuthService(authRepo, config, logger)
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

	logger.Fatal(server.Start((":" + config.Host.Port)))
}
