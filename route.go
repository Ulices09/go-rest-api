package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"
)

type Author struct {
	Name string `json:"name" validate:"required"`
}

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title" validate:"required"`
	Text   string `json:"text" validate:"required,email"`
	Author Author `json:"author" validate:"required"`
}

var posts []Post = []Post{
	{
		Id:    1,
		Title: "My post",
		Text:  "This is the text",
	},
	{
		Id:    2,
		Title: "My super post",
		Text:  "This is my super post",
	},
}

func init() {
	posts = []Post{
		{
			Id:    1,
			Title: "My post",
			Text:  "This is the text",
		},
		{
			Id:    2,
			Title: "My super post",
			Text:  "This is my super post",
		},
	}
}

func GetPosts(c echo.Context) error {
	return c.JSON(http.StatusOK, posts)
}

func GetPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	post := funk.Find(posts, func(post Post) bool {
		return post.Id == id
	})

	if post == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, post)
}

func CreatePost(c echo.Context) (err error) {
	data := new(Post)

	if err = c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(data); err != nil {
		return err
	}

	data.Id = len(posts) + 1
	posts = append(posts, *data)

	return c.JSON(http.StatusOK, data)
}
