package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoopCountsFromEnv(t *testing.T) {
	testCases := []struct {
		envValue string
		expected []int
	}{
		{"", []int{1000, 1000}},
		{"3,5,7", []int{3, 5, 7}},
		{"3,invalid,7", []int{1000, 1000}},
	}

	for _, tc := range testCases {
		t.Setenv("LOOP_COUNTS", tc.envValue)
		loopCounts := getLoopCountsFromEnv()
		assert.Equal(t, tc.expected, loopCounts)
	}
}

// Test for getDepthPrefix
func TestGetDepthPrefix(t *testing.T) {
	testCases := []struct {
		depth    int
		expected string
	}{
		{0, "a"},
		{1, "b"},
		{25, "z"},
		{26, "?"},
		{-1, "?"},
	}

	for _, tc := range testCases {
		result := getDepthPrefix(tc.depth)
		assert.Equal(t, tc.expected, result)
	}
}

func TestGenerateLoops(t *testing.T) {
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Execute the function that prints output
	loopCounts := []int{2, 2}
	generateLoops(loopCounts, 0, []string{})

	// Close the writer and read the output
	w.Close()
	var outputBuf bytes.Buffer
	n, err := outputBuf.ReadFrom(r)
	assert.Equal(t, n, int64(200))
	assert.NoError(t, err)

	// Restore the original stdout
	os.Stdout = origStdout

	// Get the actual output as string
	output := outputBuf.String()

	expectedLines := 4
	actualLines := len(strings.Split(strings.TrimSpace(output), "\n"))

	assert.Equal(t, expectedLines, actualLines)
	assert.Contains(t, output, "uuid=")
}
