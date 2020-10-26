package cacheparser

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	command := "set name ajah"
	if !Parse(command) {
		t.Log("Failed parsing")
		t.Fail()
	}
}

func TestParseFailure(t *testing.T) {
	command := "seter name ajah"
	if Parse(command) {
		t.Log("Failed parsing")
		t.Fail()
	}
}

func ExampleParse() {
	command2 := "get name"
	if Parse(command2) {
		fmt.Printf("Correct parsing")
	}
	// Output:
	// Correct parsing
}
