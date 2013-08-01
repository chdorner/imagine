package server

import (
	"net/http"
	"log"
)

func ListenAndServe(addr string) error {
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/", processErrorHandler(processHandler))

	log.Printf("server listening on %s..", addr)
	return http.ListenAndServe(addr, nil)
}
