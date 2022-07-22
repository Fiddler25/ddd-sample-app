package main

import (
	"github.com/Fiddler25/sample-app/interface/auth"
	"github.com/Fiddler25/sample-app/interface/comment"
	"github.com/Fiddler25/sample-app/interface/user"
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
	comment.Setup(root)

	e.Logger.Fatal(e.Start(":3000"))
}
