package main

import (
	"flag"
	"log"

	"github.com/cliche-niche/CS455/blob"
	"github.com/cliche-niche/CS455/util"
)

func main() {
	// Define command-line flags
	location := flag.String("location", ".", "Location to store the blob file")

	// Parse command-line arguments
	flag.Parse()

	// name and blobtype are obtained from location string
	name, blobtype, path := util.UnrollLocation(*location)

	// Initialize the Blob
	var b blob.Blob

	err := b.InitBlob(name, blobtype, path)
	if err != nil {
		log.Fatalf("Failed to initialize the blob: %v", err)
	}

	// TODO: Launch the Text Editor with contents of the blob

	defer func() {
		if err := b.CloseBlob(); err != nil {
			log.Fatalf("Failed to close the blob: %v", err)
		}
	}()

}
