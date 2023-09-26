package http

import (
	"net/http"
)

func Run(r http.Handler, addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := srv.ListenAndServe()
	return err
}
