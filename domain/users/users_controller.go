package users

import (
	"go-rest-api/types/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controller struct {
	userService UserService
}

func NewUserController(userService UserService) UserController {
	return &controller{userService}
}

func (co *controller) GetUsers(c echo.Context) error {
	users, err := co.userService.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (co *controller) GetUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := co.userService.GetById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (co *controller) CreateUser(c echo.Context) (err error) {
	data := new(entity.User)

	if err = c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(data); err != nil {
		return err
	}

	newUser, err := co.userService.Create(data)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newUser)
}
