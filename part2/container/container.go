package container

import (
	"log"

	"github.com/odk-/cme/part2/registry"
	"github.com/odk-/cme/part2/storage"
)

// RunContainer as name suggests starts a new container
// Or it will later on. For now it will only invoke download and mount of image filesystem.
func RunContainer(imageName, containerName string) error {
	registry.SetDefaultRegistry("registry-1.docker.io")
	img, err := registry.ParseImageName(imageName)
	if err != nil {
		return err
	}
	rootPath, err := storage.CreateContainerRootFS(img, containerName)
	if err != nil {
		return err
	}
	log.Println("Root path: ", rootPath)
	return nil
}
