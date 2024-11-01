package parse

import (
	"bytes"
	"encoding/json"
	"flag"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		fileName     string
		expected     User
		expectsError bool
	}{
		{
			"testdata/basic.json",
			User{"foo", 12},
			false,
		},
		{
			"testdata/nonexisting.json",
			User{},
			true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.fileName, func(t *testing.T) {
			res, err := Parse(tc.fileName)
			if err == nil && tc.expectsError {
				t.Errorf("expecting error, got success")
			}
			if err != nil && !tc.expectsError {
				t.Errorf("not expecting error, got %v", err)
			}
			if !tc.expectsError && res != tc.expected {
				t.Errorf("expecting %v, got %v", tc.expected, res)
			}
		})
	}
}

var update = flag.Bool("update", false, "update .golden.json files")

func TestParseAndIncrement(t *testing.T) {
	tests := []struct {
		fileName string
	}{
		{
			"testdata/basic.json",
		},
	}

	for _, tc := range tests {
		t.Run(tc.fileName, func(t *testing.T) {
			res, _ := ParseAndIncrementAge(tc.fileName)
			jsonRes, _ := json.Marshal(res)

			goldenFile := tc.fileName + ".golden"
			if *update {
				os.WriteFile(goldenFile, jsonRes, os.ModePerm)
			}
			expected, err := os.ReadFile(goldenFile)
			if err != nil {
				t.Errorf("failed to open golden file %s: %v", goldenFile, err)
			}
			if !bytes.Equal(expected, jsonRes) {
				t.Fail()
			}
		})
	}
}
