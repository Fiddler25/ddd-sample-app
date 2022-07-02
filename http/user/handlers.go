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

	u := user.GetUsecase(gorm.DB()).Execute(userID)

	return c.String(http.StatusOK, u.Email)
}
