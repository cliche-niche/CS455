package util

import (
	"strings"
)

// take a location path and return the name and type of the blob
func UnrollLocation(location string) (name string, blobtype string, path string) {
	ind := strings.Index(location, "/");
	if ind == -1 {
		name = location
		blobtype = strings.Split(name, ".")[1]
		path = "."
	} else {
		spl := strings.Split(location, "/")
		name = spl[len(spl)-1]
		path = strings.Join(spl[:len(spl)-1], "/")
		blobtype = strings.Split(name, ".")[1]
	}
	return name, blobtype, path
}
