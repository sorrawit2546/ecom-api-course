package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repository "github.com/sorrawit2546/internal/adapters/postgresql/sqlc"
	"github.com/sorrawit2546/internal/orders"
	"github.com/sorrawit2546/internal/products"
)

// application
type application struct {
	config config
	db     *pgx.Conn
}

// config
type config struct {
	addr string
	db   dbConfig
}

// database config
type dbConfig struct {
	dsn string
}

// mount
func (app *application) mount() http.Handler {
	//use chi
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //GET == (200)
		json.NewEncoder(w).Encode(map[string]string{"Status": "OK!"})
	})

	productService := products.NewService(*repository.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)

	orderService := orders.NewOrder(*repository.New(app.db), app.db)
	orderHandler := orders.NewOrderHandler(orderService)
	r.Post("/orders", orderHandler.PlaceOrder)

	return r
}

// run
func (app *application) run(h http.Handler) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at addr : %s", app.config.addr)
	return server.ListenAndServe()
}
