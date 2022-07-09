package user

import (
	"github.com/Fiddler25/ddd-sample-app/gorm"
	"github.com/Fiddler25/ddd-sample-app/sdk/validator"
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

	validate := validator.Validate()
	if err := validate.Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := validate.VarWithValue(req.Password, req.PasswordConfirmation, "eqfield"); err != nil {
		return c.String(http.StatusBadRequest, "パスワードが一致しません。")
	}

	res := user.NewCreateUsecase(gorm.DB()).Execute(c, req)

	return c.JSON(http.StatusCreated, res)
}

func Update(c echo.Context) error {
	var req user.UpdateRequest
	if err := c.Bind(&req); err != nil {
		log.Fatal(err)
		return err
	}

	validate := validator.Validate()
	if err := validate.Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validate.VarWithValue(req.Password, req.PasswordConfirmation, "eqfield"); err != nil {
		return c.String(http.StatusBadRequest, "パスワードが一致しません。")
	}

	return c.JSON(http.StatusCreated, "success")
}
