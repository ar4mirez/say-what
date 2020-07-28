package main_test

import (
	"testing"

	srv "github.com/ar4mirez/say-what"
)

func TestTranslate(t *testing.T) {

	response := srv.Translate("hi")
	if response != "Hola" {
		t.Fatalf("Fail TestTranslate expected value 'Hola' - returned %v", response)
	}

	response = srv.Translate("hello")
	if response != "Hola" {
		t.Fatalf("Fail TestTranslate expected value 'Hola' - returned %v", response)
	}

	response = srv.Translate("bye")
	if response != "Adios" {
		t.Fatalf("Fail TestTranslate expected value 'Adios' - returned %v", response)
	}

	response = srv.Translate("anyword")
	if response != "No implementado" {
		t.Fatalf("Fail TestTranslate expected value 'No implementado' - returned %v", response)
	}

}
