package clientauth

import (
	"strings"
	"testing"
)

func TestParseAccessTokenReturnsToken(t *testing.T) {
	token, err := parseAccessToken([]byte(`{"access_token":"abc123"}`))
	if err != nil {
		t.Fatalf("parseAccessToken returned error: %v", err)
	}
	if token != "abc123" {
		t.Fatalf("token = %q, want abc123", token)
	}
}

func TestParseAccessTokenRejectsInvalidJSON(t *testing.T) {
	_, err := parseAccessToken([]byte(`not-json`))
	if err == nil || !strings.Contains(err.Error(), "decode token response") {
		t.Fatalf("err = %v, want decode token response error", err)
	}
}

func TestParseAccessTokenRejectsMissingToken(t *testing.T) {
	_, err := parseAccessToken([]byte(`{"error":"invalid_grant"}`))
	if err == nil || !strings.Contains(err.Error(), "missing access_token") {
		t.Fatalf("err = %v, want missing access_token error", err)
	}
}

func TestParseAccessTokenRejectsNonStringToken(t *testing.T) {
	_, err := parseAccessToken([]byte(`{"access_token":42}`))
	if err == nil || !strings.Contains(err.Error(), "access_token must be a string") {
		t.Fatalf("err = %v, want access_token type error", err)
	}
}
