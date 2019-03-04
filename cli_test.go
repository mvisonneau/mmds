package main

import (
	"testing"
)

func TestRunCli(t *testing.T) {
	c := runCli()
	if c.Name != "mmds" {
		t.Fatalf("Expected c.Name to be mmds, got '%v'", c.Name)
	}
}
