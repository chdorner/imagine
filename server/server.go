package server

import (
	"github.com/chdorner/imagine/loader"
	"log"
	"net/http"
	"regexp"
)

var Version string
var originLoader *loader.Loader

func ListenAndServe(addr string, originWhitelist []*regexp.Regexp) error {
	originLoader = loader.NewLoader(originWhitelist)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/", processErrorHandler(processHandler))

	log.Printf("server listening on %s..", addr)
	return http.ListenAndServe(addr, nil)
}
