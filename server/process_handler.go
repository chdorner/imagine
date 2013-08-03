package server

import (
	"github.com/chdorner/imagine/instructions"
	"github.com/chdorner/imagine/loader"
	"github.com/chdorner/imagine/processor"
	"net/http"
	"mime"
)

func processHandler(w http.ResponseWriter, r *http.Request) {
	instr, err := instructions.ParseInstructions(r.URL.Query())
	if err != nil {
		httpErr := &httpError{err, http.StatusBadRequest}
		panic(httpErr)
	}

	p := processor.NewProcessor(instr)

	originReader, _ := loader.Load(instr.Origin)
	defer originReader.Close()

	mimetype := mime.TypeByExtension("." + instr.Format)
	w.Header().Set("Content-Type", mimetype)

	p.Process(originReader, w)
}
