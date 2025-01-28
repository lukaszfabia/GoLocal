package parsers

import (
	"backend/internal/forms"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// register forms here
type Formable interface {
	forms.Login | forms.Register |
		forms.RefreshTokenRequest | forms.EditAccount |
		forms.RestoreAccount | forms.CodeRequest | forms.NewPasswordRequest | forms.Device | forms.VoteInVotingForm
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
		return nil, errors.New("invalid JSON format")
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
	// we dont want to return error cuz image is optional
	if form == nil {
		return fileInfo, nil
	}

	fileHeaders, ok := form.File[fieldName]
	if !ok || len(fileHeaders) == 0 {
		return fileInfo, fmt.Errorf("field '%s' not found in form or no file uploaded", fieldName)
	}

	file, err := fileHeaders[0].Open()
	if err != nil {
		return fileInfo, fmt.Errorf("failed to open file: %w", err)
	}

	extension := filepath.Ext(fileHeaders[0].Filename)
	if extension == "" {
		return fileInfo, fmt.Errorf("file has no extension")
	}

	fileInfo.File = &file
	fileInfo.Extension = extension

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

/*
Transform query string to dict

Params:

  - r: Request
  - model: any - Validate json keys and insert them to dict
  - args: ...string - wanted fields

Returns:

  - list of events
  - error occured during transaction
*/
func ParseURLQuery(r *http.Request, model any, args ...string) map[string]any {
	v := reflect.ValueOf(model)
	t := reflect.TypeOf(model)

	params := map[string]any{}

	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {

			property := t.Field(i)

			tag := property.Tag.Get("json")
			if tag == "" {
				continue
			}

			param := r.URL.Query().Get(tag)
			if param == "" {
				continue
			}

			switch property.Type.Kind() {
			case reflect.Bool:
				if param == "true" {
					params[tag] = true
				} else if param == "false" {
					params[tag] = false
				}
			case reflect.Int:
				if val, err := strconv.Atoi(param); err == nil {
					params[tag] = val
				}
			case reflect.String:
				params[tag] = param
			default:
				params[tag] = param
			}
		}
	} else {
		for _, cTag := range args {
			if v := r.URL.Query().Get(cTag); v != "" {
				params[cTag] = v
			}
		}
	}

	return params
}
