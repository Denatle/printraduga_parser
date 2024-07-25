package parsers

import (
	"testing"
)

func TestCoral(t *testing.T) {
	parser := CoralTranslusentParser{}
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed with error: %v", err)
	}
	t.Logf("Result: %v", result)
}
