package loader

import (
	"io"
	"net/http"
	"regexp"
)

type Loader struct {
	whitelist []*regexp.Regexp
}

func NewLoader(whitelist []*regexp.Regexp) *Loader {
	return &Loader{whitelist}
}

func (l *Loader) Load(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
