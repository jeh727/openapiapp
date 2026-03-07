package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunServer(address string, readHeaderTimeout time.Duration) error {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	server := &http.Server{
		Addr:              address,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	log.Printf("starting server on %s", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server failed\n%w", err)
	}

	return nil
}
