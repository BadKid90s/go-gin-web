package http

import (
	"errors"
	"log"
	"net/http"
)

func Run(r http.Handler, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}
