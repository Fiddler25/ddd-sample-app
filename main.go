package main

import (
	"github.com/Fiddler25/ddd-sample-app/http/auth"
	"github.com/Fiddler25/ddd-sample-app/http/user"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	auth.Setup(e)
	user.Setup(e)

	e.Logger.Fatal(e.Start(":3000"))
}
