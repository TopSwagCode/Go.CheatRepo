package tests

import (
	"log"
	"os"
	"testing"
)

var c, python, java, dotnetcore, golang bool

// Method for setting up and tearing down tests
func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")

	c = true
	python = true
	java = false
	dotnetcore = true
	golang = true

	exitVal := m.Run()

	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestCIsAwesome(t *testing.T) {
	if c != true{
		t.Errorf("C wasn't awesome, got: %t, want: %t.", c, true)
	}
}

func TestPythonIsAwesome(t *testing.T) {
	if python != true{
		t.Errorf("Python wasn't awesome, got: %t, want: %t.", python, true)
	}
}

func TestJavaIsAwesome(t *testing.T) {
	if java != true{
		t.Errorf("Java wasn't awesome, got: %t, want: %t.", java, true)
	}
}

func TestDotnetIsCoreAwesome(t *testing.T) {
	if dotnetcore != true{
		t.Errorf("Dotnet Core wasn't awesome, got: %t, want: %t.", dotnetcore, true)
	}
}

func TestGoIsAwesome(t *testing.T) {
	if golang != true{
		t.Errorf("Java wasn't awesome, got: %t, want: %t.", golang, true)
	}
}