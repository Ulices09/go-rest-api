package users

import (
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/errors"
	httpapp "go-rest-api/internal/interface/http"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service UserService
}

func NewUserController(userService UserService) UserController {
	return &controller{userService}
}

func (co *controller) GetUsers(c echo.Context) (err error) {
	query := httpapp.GetListQuery(c)
	users, err := co.service.GetAll(query)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, users)
}

func (co *controller) GetUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError()
	}

	user, err := co.service.GetById(id)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, user)
}

func (co *controller) CreateUser(c echo.Context) (err error) {
	data := new(entity.User)

	if err = c.Bind(data); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	if err = c.Validate(data); err != nil {
		return
	}

	newUser, err := co.service.Create(data)

	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, newUser)
}
