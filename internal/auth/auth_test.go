package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	expected := "testKey"
	testHeaders := http.Header{}
	testHeaders.Add("Authorization", fmt.Sprintf("ApiKey %s", expected))

	result, err := GetAPIKey(testHeaders)

	if err != nil || result != expected {
		t.Fatalf("Expected: %s, got: %s, error: %v", expected, result, err)
	}
}
