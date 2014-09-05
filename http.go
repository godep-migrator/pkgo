package pkgo

import (
	"github.com/subosito/anoa"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

type TemplateData struct {
	User *anoa.User
}

func NewMux() *web.Mux {
	m := web.New()
	m.Get("/", HomeHandler)
	m.Get("/about", AboutHandler)
	m.Get("/logout", LogoutHandler)

	// assets
	m.Get("/robots.txt", RobotsHandler)
	m.Get("/favicon.ico", FaviconHandler)
	m.Get("/css/theme.css", CSSHandler)
	m.Get("/css/auth-icons.png", AuthIconsHandler)

	// oauth2
	m.Get("/auth", OauthHandler)
	m.Get("/auth/callback", OauthCallbackHandler)

	// middlewares
	m.Use(middleware.EnvInit)
	m.Use(AuthMiddleware)

	return m
}
