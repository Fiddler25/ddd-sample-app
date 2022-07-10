package comment

import (
	"github.com/Fiddler25/ddd-sample-app/gorm"
	"github.com/Fiddler25/ddd-sample-app/sdk/validator"
	"github.com/Fiddler25/ddd-sample-app/usecase/comment"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Create(c echo.Context) error {
	var req comment.CreateRequest
	if err := c.Bind(&req); err != nil {
		log.Fatal(err)
		return err
	}

	validate := validator.Validate()
	if err := validate.Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res := comment.NewCreateUsecase(gorm.DB()).Execute(req)

	return c.JSON(http.StatusCreated, res)
}
