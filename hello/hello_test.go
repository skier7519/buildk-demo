/*
Basic unit tests to capture the output of the example `hello.go` program
Each test case asserts that `hello.go` outputs the expected "Hello {name}!" string
A helper function was created to de-duplicate test setup, assertion, and teardown logic
*/

package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Basic test to simulate user input by setting os.Args
func TestHello(t *testing.T) {
    // Setup simulated user input
	input := "Kite Flyer"
    os.Args = []string{"hello", input}

	// Redirect stdout to temp file and perform assertion
	redirectStdoutToFileAndAssert(input, t)
}

// Test to provide an environment variable as program input 
func TestHelloWithEnvVar(t *testing.T) {
    // Setup environment variable input
	input := os.Getenv("NAME_INPUT")
    os.Args = []string{"hello", input}

	// Redirect stdout to temp file and perform assertion
	redirectStdoutToFileAndAssert(input, t)
}

// Helper function to consolidate test case logic
func redirectStdoutToFileAndAssert(input string, t *testing.T) {
    // Create a temp file
    tempFile, err := os.CreateTemp("", "output-file")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }

	// Defer cleanup of test resources
    defer os.Remove(tempFile.Name())
    defer tempFile.Close()

    // Redirect to temp file
    oldStdout := os.Stdout
    os.Stdout = tempFile

	// Run main `hello` program
    main()

    // Restore stdout
    os.Stdout = oldStdout

	// Read temp file contents
	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
	t.Fatalf("Failed to read temp file: %v", err)
	}

	// Assert temp file contains expected user input
	assert.Contains(t, string(data), "Hello, "+input+"!")
}