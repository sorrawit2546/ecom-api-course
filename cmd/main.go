package main

import (
	"context"
	"log"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/jackc/pgx/v5"
	"github.com/sorrawit2546/internal/env"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db:   dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=myuser password=mysecretpassword dbname=mydatabase sslmode=disable"),
		},
	}

	//database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	logger.Info("Connect to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db: conn,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server has failed to start, err: %s", err)
		os.Exit(1)
	}

}