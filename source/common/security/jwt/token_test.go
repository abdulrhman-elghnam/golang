package jwt

import (
	"testing"
	"time"

	"github.com/abdulrhman-elghnam/golang/source/configuration"
)

func TestGenerateAndVerifyToken(t *testing.T) {
	configuration.ACCESS_USER_TOKEN_SECRET = "test-user-secret-1234567890"
	configuration.ACCESS_USER_TOKEN_EXPIRY = time.Minute

	token, err := GenerateToken(42)
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	claims, err := VerifyToken(token, configuration.ACCESS_USER_TOKEN_SECRET)
	if err != nil {
		t.Fatalf("VerifyToken returned error: %v", err)
	}

	if claims.UserID != 42 {
		t.Fatalf("expected user ID 42, got %d", claims.UserID)
	}

	if claims.Role != UserRole {
		t.Fatalf("expected role %q, got %q", UserRole, claims.Role)
	}
}

func TestCreateLoginCredential(t *testing.T) {
	configuration.ACCESS_USER_TOKEN_SECRET = "access-secret-1234567890"
	configuration.REFRESH_USER_TOKEN_SECRET = "refresh-secret-1234567890"
	configuration.ACCESS_USER_TOKEN_EXPIRY = time.Minute
	configuration.REFRESH_TOKEN_EXPIRY = time.Hour

	credential, err := CreateLoginCredential(7, UserRole)
	if err != nil {
		t.Fatalf("CreateLoginCredential returned error: %v", err)
	}

	if credential == nil {
		t.Fatal("credential must not be nil")
	}

	if credential.AccessToken == "" || credential.RefreshToken == "" {
		t.Fatal("token strings must not be empty")
	}

	accessClaims, err := VerifyToken(credential.AccessToken, configuration.ACCESS_USER_TOKEN_SECRET)
	if err != nil {
		t.Fatalf("Verify access token error: %v", err)
	}

	if accessClaims.UserID != 7 || accessClaims.Role != UserRole {
		t.Fatal("access token claims do not match expected user role")
	}

	refreshClaims, err := VerifyToken(credential.RefreshToken, configuration.REFRESH_USER_TOKEN_SECRET)
	if err != nil {
		t.Fatalf("Verify refresh token error: %v", err)
	}

	if refreshClaims.UserID != 7 || refreshClaims.Role != UserRole {
		t.Fatal("refresh token claims do not match expected user role")
	}
}

func TestGetToken(t *testing.T) {
	got, err := GetToken("Bearer abc123")
	if err != nil {
		t.Fatalf("GetToken returned error: %v", err)
	}

	if got != "abc123" {
		t.Fatalf("expected abc123, got %q", got)
	}

	if _, err := GetToken("abc123"); err == nil {
		t.Fatal("expected malformed authorization to fail")
	}
}
