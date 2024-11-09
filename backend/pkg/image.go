package pkg

import (
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

var (
	host = fmt.Sprintf("http://localhost:%s/api/media", os.Getenv("PORT"))
)

type EventImage struct{}
type Avatar struct{}

type Saveable interface {
	GetPath(fileName string, ext string) string
}

func (e *EventImage) GetPath(fileName string, ext string) string {
	return fmt.Sprintf("%s/images/%s%s", host, fileName, ext)
}

func (a *Avatar) GetPath(fileName string, ext string) string {
	return fmt.Sprintf("%s/avatar/%s%s", host, fileName, ext)
}

// Saves image and returns new URL
func SaveImage[T Saveable](file FileInfo) (string, error) {
	var item T

	uuid := uuid.New()

	// remove old file
	if err := os.Remove(*file.OldPath); err != nil {
		return "", err
	}

	var path string = item.GetPath(uuid.String(), file.Extension)

	// create new one
	outFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, *file.File)
	if err != nil {
		return "", err
	}

	return path, nil
}
