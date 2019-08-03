package appdir

import (
	"path/filepath"
)

// AppInfo contains information required about the application
type AppInfo struct {
	Name   string
	Author string
}

// DataPath returns the absolute path for the specified path inside the app's data folder
func (info AppInfo) DataPath(path string) (string, error) {
	root, err := info.DataRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, path), nil
}

// ConfigPath returns the absolute path for the specified path inside the app's config folder
func (info AppInfo) ConfigPath(path string) (string, error) {
	root, err := info.ConfigRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, path), nil
}

// CachePath returns the absolute path for the specified path inside the app's cache folder
func (info AppInfo) CachePath(path string) (string, error) {
	root, err := info.CacheRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, path), nil
}
