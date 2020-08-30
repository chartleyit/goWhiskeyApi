package main

import (
	"os"
	"testing"
)

var startString = "Go Whiskey API backend"

func TestStart(t *testing.T) {
	if s := start(); s != startString {
		t.Error("got", start(), "expected", startString)
	}
}

func TestMain(m *testing.M) {
	eVal := m.Run()
	os.Exit(eVal)
}
