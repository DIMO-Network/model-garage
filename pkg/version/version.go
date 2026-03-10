// Package version provides the version of the application.
package version

import "golang.org/x/mod/semver"

var version = "dev"

// GetVersion returns the version of the application.
func GetVersion() string {
	cVersion := semver.Canonical(version)
	if cVersion != "" {
		return cVersion
	}
	return version
}
