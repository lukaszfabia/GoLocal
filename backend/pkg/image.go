package pkg

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

var (
	base = "./media"
	host = fmt.Sprintf("http://localhost:%s/api/media", os.Getenv("PORT"))
)

type EventImage struct{}
type Avatar struct{}

type Saveable interface {
	GetPath(base string, fileName string, ext string) string
}

func (e *EventImage) GetPath(base string, fileName string, ext string) string {
	return fmt.Sprintf("%s/images/%s%s", base, fileName, ext)
}

func (a *Avatar) GetPath(base string, fileName string, ext string) string {
	return fmt.Sprintf("%s/avatars/%s%s", base, fileName, ext)
}

// Saves image and returns new URL
func SaveImage[T Saveable](file FileInfo) (string, error) {
	var item T

	uuid := uuid.New()

	// remove old file
	if file.OldPath != nil {
		// translate path on local path
		lst := strings.Split(*file.OldPath, "/")
		filename := lst[len(lst)-1] // file.ext
		toremove := item.GetPath(base, filename, "")
		if err := os.Remove(toremove); err != nil {
			// if there was an error it means that image comes from provider
			log.Println(err)
		}
	}

	var path string = item.GetPath(base, uuid.String(), file.Extension)

	// create new one
	outFile, err := os.Create(path)
	if err != nil {
		log.Printf("Error during creating new path: %s\n", err)
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, *file.File)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var url = item.GetPath(host, uuid.String(), file.Extension)
	return url, nil
}
