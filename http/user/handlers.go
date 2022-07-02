package user

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/gorm"
	"github.com/Fiddler25/ddd-sample-app/usecase/user"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func Get(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Fatal(err)
		return err
	}

	res := user.NewGetUsecase(gorm.DB()).Execute(userID)

	return c.String(http.StatusOK, res.Email)
}

func Create(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		log.Fatal(err)
		return err
	}

	res := user.NewCreateUsecase(gorm.DB()).Execute(u)

	return c.JSON(http.StatusCreated, res)
}
