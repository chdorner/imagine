package server

import (
	"mime"
	"net/http"

	"github.com/chdorner/imagine/instructions"
	"github.com/chdorner/imagine/loader"
	"github.com/chdorner/imagine/processor"
)

func processHandler(w http.ResponseWriter, r *http.Request) {
	instr, err := instructions.ParseInstructions(r.URL.Query())
	if err != nil {
		httpErr := &httpError{err, http.StatusBadRequest}
		panic(httpErr)
	}

	p := processor.NewProcessor(instr)

	originReader, err := originLoader.Load(instr.Origin)
	if err != nil {
		var httpErr *httpError
		if serr, ok := err.(*loader.HostNotAllowedError); ok {
			httpErr = &httpError{serr, http.StatusBadRequest}
		} else if serr, ok := err.(*loader.NotFoundError); ok {
			httpErr = &httpError{serr, http.StatusNotFound}
		}
		panic(httpErr)
	}
	defer originReader.Close()

	mimetype := mime.TypeByExtension("." + instr.Format)
	w.Header().Set("Content-Type", mimetype)

	p.Process(originReader, w)
}
