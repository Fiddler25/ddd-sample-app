package main

import (
	httpUser "github.com/Fiddler25/ddd-sample-app/http/user"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	httpUser.Setup(e)

	e.Logger.Fatal(e.Start(":3000"))
}
