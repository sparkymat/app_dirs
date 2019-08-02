package appdir

import (
	"path/filepath"
)

type AppInfo struct {
	Name   string
	Author string
}

func (info AppInfo) DataPath(path string) (string, error) {
	root, err := info.DataRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, path), nil
}

func (info AppInfo) ConfigPath(path string) (string, error) {
	root, err := info.ConfigRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, path), nil
}

func (info AppInfo) CachePath(path string) (string, error) {
	root, err := info.CacheRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(root, path), nil
}
