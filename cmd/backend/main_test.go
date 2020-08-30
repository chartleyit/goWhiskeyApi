package main

import (
	"testing"
)

var startString = "Go Whiskey API backend"

func TestMain(t *testing.T) {
	if s := start(); s != startString {
		t.Error("got", start(), "expected", startString)
	}
}
