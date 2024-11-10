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
)

// register forms here
type Formable interface {
	forms.Login | forms.Register | forms.RefreshTokenRequest | forms.EditAccount | forms.RestoreAccount | forms.CodeRequest
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

func DecodeMultipartForm[T Formable](r *http.Request) (*T, error) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Println("Error parsing form:", err)
		return nil, errors.New("invalid form format")
	}

	form := new(T)

	jsonField := r.FormValue("json")
	if jsonField == "" {
		return nil, errors.New("missing json field in form")
	}

	if err := json.Unmarshal([]byte(jsonField), form); err != nil {
		log.Println("Error decoding JSON:", err)
		return nil, errors.New("invalid JSON format in form field")
	}

	return form, nil
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