package posts

import (
	"go-rest-api/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controller struct {
	postService PostService
}

func NewPostController(postService PostService) PostController {
	return &controller{postService}
}

func (co *controller) GetPosts(c echo.Context) error {
	posts, err := co.postService.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (co *controller) GetPost(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	post, err := co.postService.GetById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if post == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, post)
}

func (co *controller) CreatePost(c echo.Context) (err error) {
	data := new(entity.Post)

	if err = c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(data); err != nil {
		return err
	}

	newPost, err := co.postService.Create(data)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newPost)
}
