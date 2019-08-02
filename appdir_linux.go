// +build linux

package appdir

import (
	"os"
	"path/filepath"
)

// DataRoot returns the path to the root folder for an app's data
func (info AppInfo) DataRoot() (string, error) {
	var homeDir string
	var err error

	if homeDir, err = os.UserHomeDir(); err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".local", "share", info.Name), nil
}

// ConfigRoot returns the path to the root folder for an app's config
func (info AppInfo) ConfigRoot() (string, error) {
	var homeDir string
	var err error

	if homeDir, err = os.UserHomeDir(); err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", info.Name), nil
}

// CacheRoot returns the path to the root folder for an app's cache data
func (info AppInfo) CacheRoot() (string, error) {
	var cacheDir string
	var err error

	if cacheDir, err = os.UserCacheDir(); err != nil {
		return "", err
	}

	return filepath.Join(cacheDir, info.Name), nil
}
