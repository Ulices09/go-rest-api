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
	postService PostService
}

func NewPostController(postService PostService) PostController {
	return &controller{postService}
}

func (co *controller) GetPosts(c echo.Context) error {
	query, err := httpapp.GetListPaginatedQuery(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := co.postService.GetAll(query)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (co *controller) GetPost(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError()
	}

	post, err := co.postService.GetById(id)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, post)
}

func (co *controller) CreatePost(c echo.Context) (err error) {
	userClaims := httpapp.GetLoggedInUser(c)
	data := new(entity.Post)

	if err = c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(data); err != nil {
		return err
	}

	newPost, err := co.postService.Create(data, userClaims.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newPost)
}
