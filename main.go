package main

import (
	"github.com/Fiddler25/ddd-sample-app/http/auth"
	"github.com/Fiddler25/ddd-sample-app/http/user"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	auth.Setup(e)
	user.Setup(e)

	e.Logger.Fatal(e.Start(":3000"))
}
