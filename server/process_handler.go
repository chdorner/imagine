package server

import (
	"log"
	"net/http"
	"github.com/chdorner/imagine/instructions"
)

func processHandler(w http.ResponseWriter, r *http.Request) {
	instr, err := instructions.ParseInstructions(r.URL.Query())
	if err != nil {
		httpErr := &httpError{err, http.StatusBadRequest}
		panic(httpErr)
	}

	log.Println(instr)
}
