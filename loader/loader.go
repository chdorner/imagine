package loader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/chdorner/imagine/filecache"
)

type Loader struct {
	whitelist []*regexp.Regexp
	cache     *filecache.FileCache
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

func NewLoader(whitelist []*regexp.Regexp, cache string) *Loader {
	var fc *filecache.FileCache
	if cache != "disabled" && cache != "" {
		fc = filecache.NewFileCache(cache)
	}
	return &Loader{whitelist, fc}
}

func (l *Loader) Load(u string) (io.ReadCloser, error) {
	originUrl, err := url.ParseRequestURI(u)
	if err != nil || originUrl.Host == "" {
		return nil, &InvalidOriginError{u}
	}

	allowed := l.isAllowed(originUrl)
	if !allowed {
		return nil, &HostNotAllowedError{originUrl.Host}
	}

	if l.cache == nil {
		return l.download(u)
	}

	if l.cache.IsCached(u) {
		return l.cache.Open(u)
	} else {
		r, err := l.download(u)
		if err != nil {
			return nil, err
		}

		l.cache.Write(u, r)
		r, _ = l.cache.Open(u)
		return r, nil
	}
}

func (l *Loader) download(u string) (io.ReadCloser, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	} else if resp.StatusCode >= 400 {
		return nil, &NotFoundError{u}
	}

	return resp.Body, nil
}

func (l *Loader) isAllowed(u *url.URL) bool {
	if len(l.whitelist) == 0 {
		return true
	}

	allowed := false
	for _, r := range l.whitelist {
		if r.MatchString(u.Host) {
			allowed = true
			break
		}
	}

	return allowed
}
