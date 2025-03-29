package server

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/nedpals/supabase-go"

	"github.com/HoneySinghDev/go-echo-rest-api-template/internal/config"
	"github.com/HoneySinghDev/go-echo-rest-api-template/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	// Import postgres driver for database/sql package.
	_ "github.com/lib/pq"
)

type Router struct {
	Routes     []*echo.Route
	Root       *echo.Group
	Management *echo.Group
	APIV1Push  *echo.Group
}

type Server struct {
	Config  config.Server
	DB      *pgxpool.Pool
	Queries *db.Queries
	Echo    *echo.Echo
	SB      *supabase.Client
	Router  *Router
}

func NewServer(config config.Server) *Server {
	s := &Server{
		Config: config,
		DB:     nil,
		Echo:   nil,
		Router: nil,
	}

	return s
}

func (s *Server) Ready() bool {
	return s.DB != nil &&
		s.Echo != nil &&
		s.Router != nil
}

func (s *Server) InitDB() error {
	pgxPoolConfig, err := pgxpool.ParseConfig(s.Config.DBConnectionString())
	if err != nil {
		return err
	}
	conn, err := pgxpool.NewWithConfig(context.TODO(), pgxPoolConfig)
	if err != nil {
		return err
	}

	s.DB = conn

	queries := db.New(conn)

	s.Queries = queries

	return nil
}

func (s *Server) Start() error {
	if !s.Ready() {
		return errors.New("server is not ready")
	}

	return s.Echo.Start(s.Config.Echo.ListenAddr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Warn().Msg("Shutting down server")

	if s.DB != nil {
		log.Debug().Msg("Closing database connection")

		s.DB.Close()
	}

	log.Debug().Msg("Shutting down echo server")

	return s.Echo.Shutdown(ctx)
}
