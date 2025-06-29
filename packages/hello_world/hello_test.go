package hello_world

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintHello(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	PrintHello()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("PrintHello() output = %q; want %q", output, expected)
	}
}

func TestPrintHelloWithName(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	PrintHelloWithName("Alice")

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello, Alice!\n"
	if output != expected {
		t.Errorf("PrintHelloWithName(\"Alice\") output = %q; want %q", output, expected)
	}
}
