package logger

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestLogFormats(t *testing.T) {
	var output bytes.Buffer
	log := newLogger(&output, false)
	log.Debug("hidden")
	if output.Len() != 0 {
		t.Fatal("production debug message was logged")
	}
	log.Info("hello", "value", 42)
	var entry map[string]any
	if err := json.Unmarshal(output.Bytes(), &entry); err != nil {
		t.Fatal(err)
	}
	if entry["msg"] != "hello" || entry["value"] != float64(42) {
		t.Fatalf("unexpected log: %v", entry)
	}
	output.Reset()
	newLogger(&output, true).Debug("visible")
	if !strings.Contains(output.String(), "level=DEBUG") {
		t.Fatal("development debug message missing")
	}
}
