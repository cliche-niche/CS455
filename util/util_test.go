package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUnrollLocationWithPath(t *testing.T) {
	var location string
	var name string
	var blobtype string
	var path string

	location = "C:/Desktop/file.txt"
	
	name, blobtype, path = UnrollLocation(location)

	assert.True(t, name == "file.txt")
	assert.True(t, blobtype == "txt")
	assert.True(t, path == "C:/Desktop")
}

func TestUnrollLocationWithoutPath(t *testing.T) {
	var location string
	var name string
	var blobtype string
	var path string

	location = "file.txt"

	name, blobtype, path = UnrollLocation(location)

	assert.True(t, name == "file.txt")
	assert.True(t, blobtype == "txt")
	assert.True(t, path == ".")
}