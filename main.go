package main

import (
	"github.com/Fiddler25/ddd-sample-app/http/auth"
	"github.com/Fiddler25/ddd-sample-app/http/user"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const prefix = "/api"

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	root := e.Group(prefix)
	auth.Setup(root)
	user.Setup(root)

	e.Logger.Fatal(e.Start(":3000"))
}
