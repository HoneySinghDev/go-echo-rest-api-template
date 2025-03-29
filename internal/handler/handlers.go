package handler

import (
	"github.com/HoneySinghDev/go-echo-rest-api-template/pkg/server"
)

func AttachAllRoutes(s *server.Server) {
	// attach our routes
	AuthRoutes(s)
	DashBoardRoute(s)
}
