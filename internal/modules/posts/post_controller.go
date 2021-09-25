package posts

import (
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/errors"
	httpapp "go-rest-api/internal/interface/http"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service PostService
}

func NewPostController(postService PostService) PostController {
	return &controller{postService}
}

func (co *controller) GetPosts(c echo.Context) (err error) {
	query, err := httpapp.GetListPaginatedQuery(c)

	if err != nil {
		return errors.NewBadRequestError()
	}

	result, err := co.service.GetAll(query)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, result)
}

func (co *controller) GetPost(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError()
	}

	post, err := co.service.GetById(id)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, post)
}

func (co *controller) CreatePost(c echo.Context) (err error) {
	userClaims := httpapp.GetLoggedInUser(c)
	data := new(entity.Post)

	if err = c.Bind(data); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	if err = c.Validate(data); err != nil {
		return
	}

	newPost, err := co.service.Create(data, userClaims.ID)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, newPost)
}
