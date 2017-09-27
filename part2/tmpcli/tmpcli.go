package main

import (
	"flag"
	"log"
	"os"

	"github.com/odk-/cme/part2/container"
	"github.com/odk-/cme/part2/registry"
	"github.com/odk-/cme/part2/storage"
)

var containerName = flag.String("n", "", "name of new container [required].")
var storageRootPath = flag.String("d", "", "location of image and container files [optional].")
var imageName = flag.String("i", "", "name of image to run. Docker naming compatible [required].")
var insecureRegistry = flag.Bool("http", false, "If set registry will use http [optional].")

func main() {
	// parse flags and check if all required info was provided
	flag.Parse()
	if *containerName == "" || *imageName == "" {
		flag.Usage()
		os.Exit(1)
	}
	// initialize all packages
	if *storageRootPath != "" {
		storage.SetStorageRootPath(*storageRootPath)
	}
	registry.InsecureRegistry(*insecureRegistry)
	err := storage.InitStorage()
	if err != nil {
		log.Println(err)
	}
	// run actual container (for now it will only download an image and mount it outputing path to merged rootfs)
	err = container.RunContainer(*imageName, *containerName)
	if err != nil {
		log.Println(err)
	}
}
