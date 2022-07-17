package user

import (
	model "github.com/Fiddler25/ddd-sample-app/domain/user"
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

	if res, err := user.NewGetUsecase(gorm.DB()).Execute(model.UserID(userID)); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(http.StatusOK, res)
	}
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

	if res, err := user.NewCreateUsecase(gorm.DB()).Execute(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusCreated, res)
	}
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

	res := user.NewUpdateUsecase(gorm.DB()).Execute(req)

	return c.JSON(http.StatusNoContent, res)
}

func Delete(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Fatal(err)
		return err
	}

	user.NewDeleteUsecase(gorm.DB()).Execute(model.UserID(userID))

	return c.JSON(http.StatusNoContent, userID)
}
