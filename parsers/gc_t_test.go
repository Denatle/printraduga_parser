package parsers

import (
	"testing"
)

func TestGc(t *testing.T) {
	parser := GcTranslusentParser{}
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed with error: %v", err)
	}
	t.Logf("Result: %v", result)
}
