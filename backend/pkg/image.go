package pkg

import (
	"fmt"
	"os"
)

var (
	host = fmt.Sprintf("http://localhost:%s/api/media", os.Getenv("PORT"))
)

type EventImage struct{}
type Avatar struct{}

type Saveable interface {
	GetPath() string
}

func (e *EventImage) GetPath() string {
	return fmt.Sprintf("%s/images/", host)
}

func (a *Avatar) GetPath() string {
	return fmt.Sprintf("%s/avatar/", host)
}

func SaveImage[T Saveable](file any) error {
	// var item T

	// uuid := uuid.New()

	return nil
}
