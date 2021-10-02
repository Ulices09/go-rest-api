package posts

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/errors"
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

// GetPosts godoc
// @Summary Get posts
// @Tags posts
// @Param filter query string false "filter attributes"
// @Param page query integer false "page"
// @Param pageSize query integer false "page size"
// @Success 200 {object} dto.PaginationResult{data=[]entity.Post}
// @Failure default {object} errors.CustomError
// @Router /posts [get]
func (co *controller) GetPosts(c echo.Context) (err error) {
	query, err := dto.NewListPaginatedQuery(c)

	if err != nil {
		return errors.NewBadRequestError()
	}

	result, err := co.service.GetAll(query)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, result)
}

// GetPost godoc
// @Summary Get post
// @Tags posts
// @Param id path integer true "post id"
// @Success 200 {object} entity.Post
// @Failure default {object} errors.CustomError
// @Router /posts/{id} [get]
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

// CreatePost godoc
// @Summary Crate post
// @Tags posts
// @Security auth-token
// @Param default body posts.CreatePostRequest true "post"
// @Success 200 {object} entity.Post
// @Failure default {object} errors.CustomError
// @Router /posts [post]
func (co *controller) CreatePost(c echo.Context) (err error) {
	currentUser := entity.NewCurrentUser(c)
	data := new(CreatePostRequest)

	if err = c.Bind(data); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	if err = c.Validate(data); err != nil {
		return
	}

	newPost, err := co.service.Create(*data, currentUser.ID)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, newPost)
}
