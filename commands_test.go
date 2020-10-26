package cacheparser

import (
	"testing"
)

func TestCommands(t *testing.T) {
	_, commandExists := Commands["SET"]

	if !commandExists {
		t.Log("Command does not exist")
		t.Fail()
	}
}

func TestCommandAbsence(t *testing.T) {
	_, commandExists := Commands["HGETTER"]

	if commandExists {
		t.Log("Command should not exist")
		t.Fail()
	}
}
