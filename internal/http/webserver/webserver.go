package webserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const TIMEOUT = 30 * time.Second

type OptionServ func(s *http.Server)

func Start(port string, handler http.Handler, options ...OptionServ) error {
	srv := &http.Server{
		ReadTimeout:  TIMEOUT,
		WriteTimeout: TIMEOUT,
		Addr:         ":" + port,
		Handler:      handler,
	}

	for _, o := range options {
		o(srv)
	}

	var serverError error

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		log.Printf("Server is running on port %s", port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
			serverError = err
			sigChannel <- syscall.SIGINT
		}
		log.Println("Stopped serving new Connections")
	}()

	<-sigChannel

	log.Println("stoping server...")

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), TIMEOUT)
	defer shutdownRelease()

	err := srv.Shutdown(shutdownCtx)
	if err != nil {
		panic(err)
	}

	log.Println("Shutdown complete.")

	return serverError
}

func WithReadTimeout(t time.Duration) OptionServ {
	return func(srv *http.Server) {
		srv.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) OptionServ {
	return func(srv *http.Server) {
		srv.WriteTimeout = t
	}
}
