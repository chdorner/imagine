package server

import (
	"log"
	"net/http"
	"regexp"

	"github.com/chdorner/imagine/loader"
)

var Version string
var originLoader *loader.Loader

type Config struct {
	Addr            string
	OriginWhitelist []*regexp.Regexp
	OriginCacheDir  string
}

func ListenAndServe(c *Config) error {
	originLoader = loader.NewLoader(c.OriginWhitelist, c.OriginCacheDir)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/", processErrorHandler(processHandler))

	log.Printf("server listening on %s..", c.Addr)
	return http.ListenAndServe(c.Addr, nil)
}
