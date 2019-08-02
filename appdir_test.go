package appdir

import (
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
		assert.Equal(t, "linux", dir)
	} else if runtime.GOOS == "darwin" {
		assert.Equal(t, "darwin", dir)
	}
}
