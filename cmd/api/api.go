package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zengren707/ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.db)
	userHander := user.NewHandler(userStore)
	userHander.RegisterRoutes(subrouter)

	log.Println("listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
