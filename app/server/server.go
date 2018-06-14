package server

import (
	"net/http"

	"github.com/eliogovea/avatar/app/database"
)

type Server struct {
	db     *database.DB
	router http.ServeMux
}

func New() *Server {
	// TODO
	return nil
}
