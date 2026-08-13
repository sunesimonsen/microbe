package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	srv, err := NewServer()
	if err != nil {
		t.Fatalf("NewServer() error = %v", err)
	}

	tests := []struct {
		name             string
		path             string
		wantStatus       int
		wantLocation     string
		wantContentType  string
	}{
		{
			name:         "root redirects to about docs",
			path:         "/",
			wantStatus:   http.StatusMovedPermanently,
			wantLocation: "/docs/about",
		},
		{
			name:            "known docs page renders",
			path:            "/docs/about",
			wantStatus:      http.StatusOK,
			wantContentType: "text/html",
		},
		{
			name:       "missing docs page returns not found",
			path:       "/docs/does-not-exist",
			wantStatus: http.StatusNotFound,
		},
		{
			name:            "search route renders",
			path:            "/search?query=button",
			wantStatus:      http.StatusOK,
			wantContentType: "text/html",
		},
		{
			name:         "docs route trims trailing slash",
			path:         "/docs/about/",
			wantStatus:   http.StatusMovedPermanently,
			wantLocation: "/docs/about",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			res := httptest.NewRecorder()

			srv.ServeHTTP(res, req)

			if res.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", res.Code, tt.wantStatus)
			}

			if tt.wantLocation != "" {
				if got := res.Header().Get("Location"); got != tt.wantLocation {
					t.Fatalf("Location = %q, want %q", got, tt.wantLocation)
				}
			}

			if tt.wantContentType != "" {
				if got := res.Header().Get("Content-Type"); got != tt.wantContentType {
					t.Fatalf("Content-Type = %q, want %q", got, tt.wantContentType)
				}
			}
		})
	}
}
