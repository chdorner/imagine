package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type httpError struct {
	error
	StatusCode int
}

func logRequest(r *http.Request, statusCode int, delta time.Duration) {
	log.Printf("%s %s - %d %dms", r.Method, r.URL.String(), statusCode, delta/time.Millisecond)
}

func baseHandler(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var delta time.Duration

		defer func() {
			e := recover()
			if err, ok := e.(*httpError); ok {
				http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), err.StatusCode)
				logRequest(r, err.StatusCode, delta)
			} else if e != nil {
				logRequest(r, http.StatusInternalServerError, delta)
				panic(e)
			} else {
				logRequest(r, http.StatusOK, delta)
			}
		}()

		t := time.Now()
		f(w, r)
		delta = time.Since(t)
	}
}

func ListenAndServe(addr string) error {
	http.Handle("/status", baseHandler(statusHandler))

	log.Printf("server listening on %s..", addr)
	return http.ListenAndServe(addr, nil)
}
