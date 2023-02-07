package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// favicon handler
func favicon_handler(w http.ResponseWriter, r *http.Request) {
	// Serve the favicon
	http.ServeFile(w, r, "./assets/art/favicon.ico")
}

// teapot handler
func teapot_handler(w http.ResponseWriter, r *http.Request) {
	// return teapot state
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusTeapot)
	// serve a teapot image with linkage
	log.Print("Tea Served.")
	io.WriteString(w, "<html><h1><a href='https://datatracker.ietf.org/doc/html/rfc2324/'>HTCPTP</h1><img src='https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftaooftea.com%2Fwp-content%2Fuploads%2F2015%2F12%2Fyixing-dark-brown-small.jpg&f=1&nofb=1' alt='Im a teapot'></a><html>")
}

// Main server program
func main() {
	// initialize a time.Duration variable to hold a wait time-period
	var wait time.Duration
	// define a graceful termination period
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Initialize a new router
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")

	srv := &http.Server{
		// address to listen on
		Addr: "0.0.0.0:" + PORT,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		// router to serve as my handler
		Handler: router, // Pass our instance of gorilla/mux in.
	}

	// site icon server
	router.HandleFunc("/favicon.ico", favicon_handler).Methods("GET")
	// teapot handler
	router.HandleFunc("/teapot", teapot_handler).Methods("GET")
	// define the fileserver root dir
	router.PathPrefix("/").Handler(http.StripPrefix("/",
		http.FileServer(http.Dir("../Production"))))

	// pass all requests to my router
	http.Handle("/", router)
	// print listener status
	log.Print("Listening at http://localhost:8080")

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// define channel to recieve server shutdown signal
	shutdown_chan := make(chan os.Signal, 1)
	// I'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(shutdown_chan, os.Interrupt)

	// Block until I receive the signal in channel c
	<-shutdown_chan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, I could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if my application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	// exit the program succesfully
	os.Exit(0)
}
