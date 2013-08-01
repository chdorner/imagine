package server

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type RequestInstructions struct {
	Origin string
	Action string
	Format string
	Width  int
	Height int
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	instructions := parseInstructions(r.URL.Query())
	log.Println(instructions)
}

func parseInstructions(p url.Values) *RequestInstructions {
	var err error

	i := &RequestInstructions{}
	i.Origin = p.Get("origin")
	i.Action = p.Get("action")
	i.Format = p.Get("format")

	i.Width, err = strconv.Atoi(p.Get("width"))
	if err != nil {
		err = &httpError{errors.New("width is not an integer"), http.StatusBadRequest}
		panic(err)
	}

	i.Height, err = strconv.Atoi(p.Get("height"))
	if err != nil {
		err = &httpError{errors.New("height is not an integer"), http.StatusBadRequest}
		panic(err)
	}

	return i
}
