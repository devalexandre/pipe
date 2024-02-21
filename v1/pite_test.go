package v1

import (
	"strings"
	"testing"
)

func TestPipe(t *testing.T) {

	pipeline := Pipe(
		strings.ToUpper,
		strings.TrimSpace,
	)

	input := "   go is awesome   "
	expected := "GO IS AWESOME"
	result, err := pipeline(input)
	if err != nil {
		t.Errorf("Pipeline returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestParseTo(t *testing.T) {

	result := "test result"
	var target string

	err := ParseTo(result, &target)
	if err != nil {
		t.Errorf("ParseTo returned an error: %v", err)
	}

	if target != result {
		t.Errorf("Expected target to be %s, got %s", result, target)
	}

	var intTarget int
	err = ParseTo(result, &intTarget)
	if err == nil {
		t.Errorf("Expected ParseTo to fail when types are incompatible")
	}
}
