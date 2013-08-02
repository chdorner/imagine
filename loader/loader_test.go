package loader

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"bytes"
)

func TestLoad(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	testDir := filepath.Join(pwd, "..", "test")
	srv := httptest.NewServer(http.FileServer(http.Dir(testDir)))
	defer srv.Close()

	expected, err := ioutil.ReadFile(filepath.Join(testDir, "image.jpg"))
	if err != nil {
		t.Fatal(err)
	}

	reader, err := Load(srv.URL + "/image.jpg")
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
