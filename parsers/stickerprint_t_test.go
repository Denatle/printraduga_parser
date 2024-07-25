package parsers

import (
	"testing"
)

func TestStickerprint(t *testing.T) {
	parser := StickerPrintTranslusentParser{}
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Failed with error: %v", err)
	}
	t.Logf("Result: %v", result)
}
