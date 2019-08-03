package appdir

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	appInfo = AppInfo{Name: "TestApp", Author: "TestAuthor"}
)

func TestDataDir(t *testing.T) {
	dir, err := appInfo.DataRoot()
	assert.Nil(t, err)

	if runtime.GOOS == "windows" {
		assert.Equal(t, "windows", dir)
	} else if runtime.GOOS == "linux" {
		homeDir := os.Getenv("HOME")
		assert.Equal(t, fmt.Sprintf("%v/.local/share/TestApp", homeDir), dir)
	} else if runtime.GOOS == "darwin" {
		assert.Equal(t, "darwin", dir)
	}
}

func TestConfigDir(t *testing.T) {
	dir, err := appInfo.ConfigRoot()
	assert.Nil(t, err)

	if runtime.GOOS == "windows" {
		assert.Equal(t, "windows", dir)
	} else if runtime.GOOS == "linux" {
		homeDir := os.Getenv("HOME")
		assert.Equal(t, fmt.Sprintf("%v/.config/TestApp", homeDir), dir)
	} else if runtime.GOOS == "darwin" {
		assert.Equal(t, "darwin", dir)
	}
}

func TestCacheDir(t *testing.T) {
	dir, err := appInfo.CacheRoot()
	assert.Nil(t, err)

	if runtime.GOOS == "windows" {
		assert.Equal(t, "windows", dir)
	} else if runtime.GOOS == "linux" {
		homeDir := os.Getenv("HOME")
		assert.Equal(t, fmt.Sprintf("%v/.cache/TestApp", homeDir), dir)
	} else if runtime.GOOS == "darwin" {
		assert.Equal(t, "darwin", dir)
	}
}
