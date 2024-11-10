package store

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

type Store interface {
	SetCode(email string) (string, error)
	Compare(email, code string) bool
	clear(email string)
}

// Use cases: verification codes, password restoring
type storage struct {
	sync.Mutex
	usersCodes map[string]string // map: email - code
}

func New() *storage {
	log.Println("Store has been initialized")
	return &storage{
		usersCodes: map[string]string{},
	}
}

// Removes row with email and code
func (s *storage) clear(email string) {
	s.Lock()
	defer s.Unlock()
	delete(s.usersCodes, email)
}

/*
	Sets generated 6 number code for user

Params:

  - email: requested users email

Returns:

  - code
  - error occured during process
*/
func (s *storage) SetCode(email string) (string, error) {
	s.Lock()
	defer s.Unlock()

	code, err := generateCode(6)
	if err != nil {
		return "", err
	}

	s.usersCodes[email] = code
	go s.setExp(time.Minute, email)

	return code, nil
}

// Sets expiration of a code
func (s *storage) setExp(t time.Duration, email string) {
	time.Sleep(t) // sleep for t and then clear
	s.clear(email)
}

// Comapres incoming code with stored
func (s *storage) Compare(email, code string) bool {
	s.Lock()
	defer s.Unlock()
	c, ok := s.usersCodes[email]

	// is email exists
	if !ok {
		return false
	}

	// compare codes
	return c == code
}

func generateCode(len int) (string, error) {
	if len < 6 {
		return "", errors.New("given length is not safe")
	}

	rand.Seed(uint64(time.Now().UnixNano()))

	var buf bytes.Buffer

	for i := 0; i < len; i++ {
		n := fmt.Sprint(rand.Intn(10))
		buf.WriteString(n)

	}

	return buf.String(), nil
}
