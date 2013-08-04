package filecache

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func tempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "imagine-")
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func touchFile(t *testing.T, path, content string) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(path, []byte(content), 0664); err != nil {
		t.Fatal(err)
	}
}

func TestIsCached(t *testing.T) {
	root := tempDir(t)
	defer os.RemoveAll(root)
	fc := NewFileCache(root)

	uri := "http://example.com/test.png"
	path := fc.path(uri)

	if fc.IsCached(uri) {
		t.Fatal("cached file should not exist")
	}

	touchFile(t, path, "")

	if !fc.IsCached(uri) {
		t.Fatal("cached file should be there")
	}
}

func TestOpen(t *testing.T) {
	root := tempDir(t)
	defer os.RemoveAll(root)
	fc := NewFileCache(root)

	uri := "http://example.com/test.png"
	path := fc.path(uri)

	_, err := fc.Open(uri)
	if err == nil {
		t.Fatal("open should have returned an error, cache file shouldn't exist")
	}

	content := "Hello"
	touchFile(t, path, content)

	f, err := fc.Open(uri)
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	actual := string(data)
	if actual != content {
		t.Fatal("open seemed to have opened a wrong file, got content: %s", actual)
	}
}

func TestWrite(t *testing.T) {
	root := tempDir(t)
	//defer os.RemoveAll(root)
	fc := NewFileCache(root)

	uri := "http://example.com/test.png"
	path := fc.path(uri)
	content := "the content"
	r := strings.NewReader(content)

	err := fc.Write(uri, r)
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	actual := string(data)
	if actual != content {
		t.Fatalf("write seemed to have written the wrong content, got: %s", actual)
	}
}

func TestPath(t *testing.T) {
	root := tempDir(t)
	defer os.RemoveAll(root)
	fc := NewFileCache(root)

	actual := fc.path("http://example.com/test.png")
	expected := root + "/9c/d8/f4/1a/4a/6a/04/94/7f/48/ce/45/49/b6/82/8d.png"

	if expected != actual {
		t.Fatalf("path is not as expected, got: %s", actual)
	}
}
