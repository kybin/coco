package main

import (
	"testing"
)

func TestParseJobFile(t *testing.T) {
	j, err := parseJobFile("testdata/sample.job")
	if err != nil {
		t.Fatalf("could not parse job file: %v", err)
	}
	if j == nil {
		t.Fatalf("could not parse job file: nil job")
	}
}
