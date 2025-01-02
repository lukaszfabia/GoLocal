package store_test

import (
	"backend/internal/store"
	"testing"
)

func TestSingleSave(t *testing.T) {
	store := store.New()
	email0 := "email"

	code0, err0 := store.SetCode(email0)

	if err0 != nil {
		t.Errorf(err0.Error())
	}

	if !store.Compare(email0, code0) {
		t.Error("Something went wrong during saving code")
	}
}

func TestDoubleSave(t *testing.T) {
	store := store.New()
	email0 := "email0"
	email1 := "email1"

	code0, err0 := store.SetCode(email0)
	code1, err1 := store.SetCode(email1)

	if err0 != nil {
		t.Errorf(err0.Error())
	}
	if err1 != nil {
		t.Errorf(err1.Error())
	}

	if !store.Compare(email1, code1) {
		t.Error("Something went wrong during saving code")
	}

	if !store.Compare(email0, code0) {
		t.Error("Something went wrong during saving code")
	}
}

func TestOverwriteCode(t *testing.T) {
	store := store.New()
	email0 := "email0"
	email1 := email0

	code0, err0 := store.SetCode(email0)
	//there should overwrite code
	code1, err1 := store.SetCode(email1)

	if err0 != nil {
		t.Errorf(err0.Error())
	}
	if err1 != nil {
		t.Errorf(err1.Error())
	}

	// expected true
	if !store.Compare(email1, code1) {
		t.Error("Something went wrong during saving code")
	}

	// expected false
	if store.Compare(email0, code0) {
		t.Errorf("Something went wrong during overwriting code")
	}
}
