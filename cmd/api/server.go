package main  

import (
	"fmt" 
	"log" //Importing the log package for logging errors and information.
	"net/http" //Importing the net/http package for handling HTTP requests and responses.
	"time"
)
// 
func(app *application)  serve() error {
	server := &http.Server{
		Addr: fmt.Sprintf(":&d", app.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 *time.Second,

	}
	log.Printf("Starting server on port %d", app.port)

	return server.ListenAndServe()
}