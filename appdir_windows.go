// +build windows

package appdir

import (
	"errors"
	"os"
	"path/filepath"
)

// DataRoot returns the path to the root folder for an app's data
func (info AppInfo) DataRoot() (string, error) {
	var appDataDir string

	if appDataDir = os.Getenv("APPDATA"); appDataDir == "" {
		return "", errors.New("%APPDATA% not set")
	}

	return filepath.Join(appDataDir, "Roaming", info.Author, info.Name), nil
}

// ConfigRoot returns the path to the root folder for an app's config
func (info AppInfo) ConfigRoot() (string, error) {
	// In Windows, there are no separate config and data directories
	return info.DataRoot()
}

// CacheRoot returns the path to the root folder for an app's cache data
func (info AppInfo) CacheRoot() (string, error) {
	var appDataDir string

	if appDataDir = os.Getenv("APPDATA"); appDataDir == "" {
		return "", errors.New("%APPDATA% not set")
	}

	return filepath.Join(appDataDir, "Local", info.Author, info.Name), nil
}
