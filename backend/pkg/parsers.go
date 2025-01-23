package pkg

import (
	"backend/internal/forms"
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// register forms here
type Formable interface {
	forms.Login | forms.Register |
		forms.RefreshTokenRequest | forms.EditAccount |
		forms.RestoreAccount | forms.CodeRequest | forms.NewPasswordRequest | forms.Device
}

type FileInfo struct {
	File      *multipart.File
	Extension string
	OldPath   *string
}

func DecodeJSON[T Formable](r *http.Request) (*T, error) {
	form := new(T) // new instance of T
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Println("Error decoding JSON:", err)
		return nil, errors.New("Invalid JSON format")
	}

	return form, nil
}

func DecodeMultipartForm[T any](r *http.Request) (*T, error) {
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
		log.Println("Error parsing form:", err)
		return nil, errors.New("invalid form format")
	}

	form := new(T)

	// map values
	if err := decoder.Decode(form, r.PostForm); err != nil {
		log.Println("Error decoding form fields:", err)
		return nil, errors.New("invalid form data")
	}

	return form, nil
}

func ParseDate(date string) time.Time {
	if date, err := time.Parse(time.DateOnly, date); err != nil {
		log.Println("Failed to parse date")
		return time.Time{}
	} else {
		return date
	}
}

func GetFileFromForm(form *multipart.Form, fieldName string) (FileInfo, error) {
	fileInfo := FileInfo{
		File:      nil,
		Extension: "",
	}

	fileHeaders := form.File[fieldName]
	if len(fileHeaders) == 0 {
		return fileInfo, errors.New("file not found in form")
	}

	file, err := fileHeaders[0].Open() // one file
	if err != nil {
		return fileInfo, errors.New("failed to retrieve file")
	}
	defer file.Close()

	fileInfo.File = &file

	fileInfo.Extension = filepath.Ext(fileHeaders[0].Filename)

	return fileInfo, nil
}

func ParseHTMLToString(templateName string, data any) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %v", err)
		return "", errors.New("failed to get working directory")
	}
	log.Println(pwd)

	templatePath := filepath.Join(pwd, "templates", templateName)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Error parsing template file %s: %v", templateName, err)
		return "", errors.New("failed to parse email template")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", templateName, err)
		return "", errors.New("failed to execute email template")
	}

	body := buf.String()
	if body == "" {
		log.Println("Email body is empty")
		return "", errors.New("email body is empty")
	}

	return body, nil
}

func ParseURLQuery(r *http.Request, model any, args ...string) map[string]any {
	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model)

	// map with tags and values for q
	params := map[string]any{}

	if v.Kind() == reflect.Struct {
		// iter through all event fields
		for i := 0; i < v.NumField(); i++ {
			// for struct
			property := t.Field(i)
			tag := property.Tag.Get("json") // get json tag
			// skip field without json tag
			if tag == "" {
				continue
			}

			// skip field without param
			param := r.URL.Query().Get(tag)
			if param == "" {
				continue
			}

			params[tag] = param
		}
	}

	// interate through customtags
	// add additional params
	for _, cTag := range args {
		if v := r.URL.Query().Get(cTag); v != "" {
			params[cTag] = v
		}
	}

	return params
}
