package harness_test

import (
	"os"
	"strings"
	"testing"

	"github.com/px0-ai/harness"
)

func TestVersion(t *testing.T) {
	data, err := os.ReadFile("VERSION")
	if err != nil {
		t.Fatalf("failed to read VERSION: %v", err)
	}
	expected := strings.TrimSpace(string(data))
	if harness.Version != expected {
		t.Fatalf("Version = %q, want %q", harness.Version, expected)
	}
	if harness.Version == "" {
		t.Fatal("Version is empty")
	}
}
