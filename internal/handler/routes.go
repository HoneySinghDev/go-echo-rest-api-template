package handler

import (
	"github.com/HoneySinghDev/go-echo-rest-api-template/internal/handler/auth"
	"github.com/HoneySinghDev/go-echo-rest-api-template/internal/handler/dashboard"
	"github.com/HoneySinghDev/go-echo-rest-api-template/internal/middleware"
	"github.com/HoneySinghDev/go-echo-rest-api-template/pkg/server"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(s *server.Server) {
	a := s.Echo.Group("/api/auth")

	routes := []*echo.Route{
		// Auth - Login
		a.POST("/login", auth.HandleLoginCreate(s)),
		// Auth - Signup
		a.POST("/signup", auth.HandleSignupCreate(s)),
	}

	s.Router.Routes = append(s.Router.Routes, routes...)
}

func DashBoardRoute(s *server.Server) {
	d := s.Echo.Group("/api/dashboard", middleware.WithAuth)

	routes := []*echo.Route{
		d.GET("", dashboard.HandleDashboard(s)),
	}

	s.Router.Routes = append(s.Router.Routes, routes...)
}
