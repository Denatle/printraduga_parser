package parsers

import (
	"testing"
)

func TestDigital(t *testing.T) {
	parser := DigitalTranslusentParser{}
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed with error: %v", err)
	}
	t.Logf("Result: %v", result)
}
