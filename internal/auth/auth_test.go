package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "valid ApiKey header",
			header:  http.Header{"Authorization": []string{"ApiKey abc123"}},
			wantKey: "abc123",
			wantErr: false,
		},
		{
			name:    "missing header",
			header:  http.Header{},
			wantErr: true,
		},
		{
			name:    "wrong scheme",
			header:  http.Header{"Authorization": []string{"Bearer abc123"}},
			wantErr: true,
		},
		{
			name:    "too many parts",
			header:  http.Header{"Authorization": []string{"ApiKey abc 123"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got key %q", key)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if key != tt.wantKey {
				t.Fatalf("got key %q, want %q", key, tt.wantKey)
			}
		})
	}
}
