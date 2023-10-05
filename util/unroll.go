package util

import (
	"fmt"
	"strings"
)

// take a location path and return the name and type of the blob
func UnrollLocation(location string) (name string, blobtype string, path string) {
	fmt.Println(location)
	spl := strings.Split(location, "/")
	name = spl[len(spl)-1]
	path = strings.Join(spl[:len(spl)-1], "/")
	blobtype = strings.Split(name, ".")[1]
	fmt.Println(name, blobtype)
	return name, blobtype, path
}
