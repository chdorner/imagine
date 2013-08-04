package loader

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"regexp"
	"testing"
)

func TestLoad(t *testing.T) {
	testDir := "../test"
	srv := httptest.NewServer(http.FileServer(http.Dir(testDir)))
	defer srv.Close()

	expected, err := ioutil.ReadFile(filepath.Join(testDir, "rectangle.jpg"))
	if err != nil {
		t.Fatal(err)
	}

	originLoader := NewLoader([]*regexp.Regexp{regexp.MustCompile(`.*`)}, "disabled")

	reader, err := originLoader.Load(srv.URL + "/rectangle.jpg")
	defer reader.Close()
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Fatal("expected different data when loading image.jpg")
	}
}

func TestLoadInvalidUrl(t *testing.T) {
	l := NewLoader([]*regexp.Regexp{regexp.MustCompile(`.*`)}, "disabled")

	_, err := l.Load("test.png")
	if _, ok := err.(*InvalidOriginError); !ok {
		t.Fatal("should have returned InvalidOriginError")
	}

	_, err = l.Load("/test.png")
	if _, ok := err.(*InvalidOriginError); !ok {
		t.Fatal("should have returned InvalidOriginError")
	}
}

func TestLoadNotAllowed(t *testing.T) {
	l := NewLoader([]*regexp.Regexp{regexp.MustCompile(`example\.com`)}, "disabled")

	_, err := l.Load("http://example2.com/test.png")
	if _, ok := err.(*HostNotAllowedError); !ok {
		t.Fatal("should have returned HostNotAllowedError")
	}
}

func TestNotFound(t *testing.T) {
	testDir := "../test"
	srv := httptest.NewServer(http.FileServer(http.Dir(testDir)))
	defer srv.Close()

	l := NewLoader([]*regexp.Regexp{regexp.MustCompile(`.*`)}, "disabled")

	_, err := l.Load(srv.URL + "/notfound.png")
	if _, ok := err.(*NotFoundError); !ok {
		t.Fatal("should have returned NotFoundError")
	}
}
