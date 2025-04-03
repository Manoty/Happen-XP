package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Manoty/Happen-XP/internal/database"
	"github.com/Manoty/Happen-XP/internal/env"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)
type application struct {
	port      int
	jwtSecret string
	models    database.Models
}
func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port: env.GetEnvInt("PORT", 9090),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-12345678"),
		models: models,
	}
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

// serve method starts the HTTP server and listens for incoming requests.
func (app *application) serve() error {
	addr := fmt.Sprintf(":%d", app.port)
	log.Printf("starting server on %s", addr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running"))
	})

	return http.ListenAndServe(addr, nil)

}