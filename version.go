package patentfetch

import (
	"strings"
)

const (
    // Version number of release
    Version = "0.0.1"

    // ReleaseDate, the date version.go was generated
    ReleaseDate = "2025-10-17"

    // ReleaseHash, the Git hash when version.go was generated
    ReleaseHash = ""

)

// FmtHelp lets you process a text block with simple curly brace markup.
func FmtHelp(src string, appName string, version string, releaseDate string, releaseHash string) string {
	m := map[string]string {
		"{app_name}": appName,
		"{version}": version,
		"{release_date}": releaseDate,
		"{release_hash}": releaseHash,
	}
	for k, v := range m {
		if strings.Contains(src, k) {
			src = strings.ReplaceAll(src, k, v)
		}
	}
	return src
}

