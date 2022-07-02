package user

import (
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
	var req user.CreateRequest
	if err := c.Bind(&req); err != nil {
		log.Fatal(err)
		return err
	}
	if req.Password != req.PasswordConfirmation {
		return c.String(http.StatusBadRequest, "パスワードが一致しません。")
	}

	res := user.NewCreateUsecase(gorm.DB()).Execute(req)

	return c.JSON(http.StatusCreated, res)
}
