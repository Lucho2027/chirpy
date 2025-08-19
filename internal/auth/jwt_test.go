package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	tests := []struct {
		name     string
		userID   uuid.UUID
		secret   string
		duration time.Duration
		wantErr  bool
	}{
		{
			name:     "Valid JWT Creation",
			userID:   uuid.New(),
			secret:   "mysecret",
			duration: time.Minute,
			wantErr:  false,
		},
		{
			name:     "Empty Secret",
			userID:   uuid.New(),
			secret:   "",
			duration: time.Minute,
			wantErr:  true,
		},
		{
			name:     "Empty duration",
			userID:   uuid.New(),
			secret:   "mysecret",
			duration: 0,
			wantErr:  true,
		},
		{
			name:     "Empty UUID",
			userID:   uuid.Nil,
			secret:   "mysecret",
			duration: time.Minute,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := MakeJWT(tt.userID, tt.secret, tt.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && token == "" {
				t.Error("MakeJWT() returned empty token for valid input")
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	// Create a valid user ID and token for testing
	userID := uuid.New()
	secret := "test-secret"
	validToken, err := MakeJWT(userID, secret, time.Minute)

	if err != nil {
		t.Fatalf("Failed to create test token: %v", err)
	}

	tests := []struct {
		name        string
		tokenString string
		tokenSecret string
		wantUserID  uuid.UUID
		wantErr     bool
	}{
		{
			name:        "Valid Token",
			tokenString: validToken,
			tokenSecret: secret,
			wantUserID:  userID,
			wantErr:     false,
		},
		{
			name:        "Ivalid Token",
			tokenString: validToken,
			tokenSecret: "wrong-secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		}, {
			name:        "Empty Token",
			tokenString: "",
			tokenSecret: secret,
			wantUserID:  uuid.Nil,
			wantErr:     true,
		}, {
			name:        "Malformed Token",
			tokenString: "waka.waka.ee",
			tokenSecret: secret,
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserID, err := ValidateJWT(tt.tokenString, tt.tokenSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserID != tt.wantUserID {
				t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
			}
		})
	}
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name        string
		setupHeader func() http.Header
		wantErr     bool
		wantToken   string
	}{
		{
			name: "Bearer is present",
			setupHeader: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Bearer this-be-the-token")
				return h
			},
			wantErr:   false,
			wantToken: "this-be-the-token",
		},
		{
			name: "Missing Auth Token",
			setupHeader: func() http.Header {
				h := http.Header{}
				h.Set("", "")
				return h
			},
			wantErr:   true,
			wantToken: "",
		},
		{
			name: "Empty Auth Header",
			setupHeader: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "")
				return h
			},
			wantErr:   true,
			wantToken: "",
		}, {
			name: "Auth W/o Bearer prefix",
			setupHeader: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "this-be-the-token")
				return h
			},
			wantErr:   true,
			wantToken: "",
		}, {
			name: "Auth w/prefix not token",
			setupHeader: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Bearer ")
				return h
			},
			wantErr:   true,
			wantToken: "",
		},
		{
			name: "different  auth",
			setupHeader: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Basic this-be-the-token")
				return h
			},
			wantErr:   true,
			wantToken: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := tt.setupHeader()
			token, err := GetBearerToken(headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && token != tt.wantToken {
				t.Errorf("GetBearerToken() token = %v, want %v", token, tt.wantToken)
			}
		})
	}
}
