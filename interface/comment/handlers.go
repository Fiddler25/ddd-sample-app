package comment

import (
	"github.com/Fiddler25/sample-app/db"
	"github.com/Fiddler25/sample-app/sdk/validator"
	"github.com/Fiddler25/sample-app/usecase/comment"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Create(c echo.Context) error {
	var req comment.CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	validate := validator.Validate()
	if err := validate.Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res := comment.NewCreateUsecase(db.Conn()).Execute(req)

	return c.JSON(http.StatusCreated, res)
}
