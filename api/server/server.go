package server

import (
	"database/sql"
	"fmt"

	"github.com/anneau/go-template/api"
	"github.com/anneau/go-template/config"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	engine *gin.Engine
	config *config.HTTPServerConfig
}

func NewServer(config *config.HTTPServerConfig, dbConn *sql.DB) (*HTTPServer, error) {
	engine := gin.New()

	r, err := api.InitializeRouter()

	if err != nil {
		return nil, err
	}

	r.Setup(&engine.RouterGroup)

	return &HTTPServer{
		engine: engine,
		config: config,
	}, nil
}

func (s *HTTPServer) Run() error {
	return s.engine.Run(fmt.Sprintf(":%d", s.config.Port))
}
