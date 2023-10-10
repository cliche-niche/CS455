package main

import (
	"flag"
	"log"
	"os"

	"github.com/cliche-niche/CS455/blob"
	"github.com/cliche-niche/CS455/cli"
	"github.com/cliche-niche/CS455/util"
)

func main() {
	// Define command-line flags
	location := flag.String("location", ".", "Location to store the blob file")
	isFile := true

	// Parse command-line arguments
	flag.Parse()

	info, err := os.Stat(*location)
	if err != nil {
		panic(err)
	}

	if info.IsDir() {
		isFile = false
	}
	
	// Initialize and run the cli application
	var cli cli.Cli

	if isFile {
		// name and blobtype are obtained from location string
		name, blobtype, path := util.UnrollLocation(*location)

		// Initialize the Blob
		var b blob.Blob

		err := b.InitBlob(name, blobtype, path)
		if err != nil {
			log.Fatalf("Failed to initialize the blob: %v", err)
		}

		cli.InitCli(&b, "", false)
	} else {
		cli.InitCli(nil, *location, true)
		cli.FileChangeModal()
	}

	cli.AddExitModal()
	cli.AppInputCapture()

	err = cli.RunApp()
	if err != nil {
		panic(err)
	}
}
