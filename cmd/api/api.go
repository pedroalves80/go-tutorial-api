package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedroalves80/go-tutorial-api/service/cart"
	"github.com/pedroalves80/go-tutorial-api/service/order"
	"github.com/pedroalves80/go-tutorial-api/service/product"
	"github.com/pedroalves80/go-tutorial-api/service/user"
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
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandle := cart.NewHandler(productStore, orderStore, userStore)
	cartHandle.RegisterRoutes(subrouter)

	log.Println("Listening on: ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
