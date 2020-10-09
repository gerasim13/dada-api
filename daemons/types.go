package daemons

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	DefaultLifeSpan     = 30 * time.Second
	DefaultLifeSpanHalf = DefaultLifeSpan / 2
)

var (
	httpClient = http.Client{Timeout: 30 * time.Second}
)

type Daemon interface {
	Run(ctx context.Context)
	GetData() interface{}
}

type storage interface {
	save(bs []byte) error
	isStorageValid() bool
}

type fileStorage struct {
	FilePath string
	LifeSpan time.Duration
}

type Logger interface {
	Error(e ...interface{})
}

func (fs *fileStorage) read() ([]byte, error) {
	return ioutil.ReadFile(fs.FilePath)
}

func (fs *fileStorage) save(bs []byte) error {
	return ioutil.WriteFile(filepath.Clean(fs.FilePath), bs, 0600)
}

func (fs *fileStorage) isStorageValid() bool {
	fileInfo, err := os.Stat(fs.FilePath)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	if fileInfo.ModTime().Before(time.Now().Add(-1 * fs.LifeSpan)) {
		return false
	}
	return true
}
