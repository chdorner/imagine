package filecache

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileCache struct {
	root string
}

func NewFileCache(root string) *FileCache {
	return &FileCache{root}
}

func (c *FileCache) IsCached(uri string) bool {
	stat, err := os.Stat(c.path(uri))
	if err != nil {
		return false
	}

	if stat.IsDir() {
		return false
	}

	return true
}

func (c *FileCache) Open(uri string) (io.ReadCloser, error) {
	f, err := os.Open(c.path(uri))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (c *FileCache) Write(uri string, r io.Reader) error {
	path := c.path(uri)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}

	return nil
}

func (c *FileCache) path(uri string) string {
	path := c.root

	h := md5.New()
	io.WriteString(h, uri)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	r := regexp.MustCompile(`(.{2})`)
	matches := r.FindAllString(hash, -1)

	path += "/" + strings.Join(matches, "/")

	path += filepath.Ext(uri)

	return path
}
