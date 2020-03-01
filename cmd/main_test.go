package main

import (
	"os"
	"testing"
)

func TestGetSecret(t *testing.T) {
	res := os.Getenv("OLD_SECRET")
	if "" == res {
		t.Errorf("OLD_SECRET = <empty>; want <not empty>")
	} else {
		t.Log("OLD_SECRET = " + res)
	}
}
