package loader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

type Loader struct {
	whitelist []*regexp.Regexp
}

type HostNotAllowedError struct {
	Host string
}

func (e HostNotAllowedError) Error() string {
	return fmt.Sprintf("host %s is not allowed", e.Host)
}

type NotFoundError struct {
	Url string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("origin %s could not be found", e.Url)
}

type InvalidOriginError struct {
	Url string
}

func (e InvalidOriginError) Error() string {
	return fmt.Sprintf("origin %s is not a valid url", e.Url)
}

func NewLoader(whitelist []*regexp.Regexp) *Loader {
	return &Loader{whitelist}
}

func (l *Loader) Load(u string) (io.ReadCloser, error) {
	originUrl, err := url.ParseRequestURI(u)
	if err != nil || originUrl.Host == "" {
		return nil, &InvalidOriginError{u}
	}

	allowed := false
	for _, r := range l.whitelist {
		if r.MatchString(originUrl.Host) {
			allowed = true
			break
		}
	}
	if !allowed {
		return nil, &HostNotAllowedError{originUrl.Host}
	}

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	} else if resp.StatusCode >= 400 {
		return nil, &NotFoundError{u}
	}

	return resp.Body, nil
}
