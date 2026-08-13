package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
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
		referer          string
		wantStatus       int
		wantLocation     string
		wantContentType  string
		wantBodyContains []string
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
			name:             "search route renders without results",
			path:             "/search?query=this-should-never-match",
			wantStatus:       http.StatusOK,
			wantContentType:  "text/html",
			wantBodyContains: []string{"No results for", "this-should-never-match"},
		},
		{
			name:            "search route ignores malformed referer",
			path:            "/search?query=button",
			referer:         ":bad-url",
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
				if tt.referer != "" {
					req.Header.Set("Referer", tt.referer)
				}
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

				if len(tt.wantBodyContains) > 0 {
					body := res.Body.String()
					for _, want := range tt.wantBodyContains {
						if !strings.Contains(body, want) {
							t.Fatalf("body = %q, want substring %q", body, want)
						}
					}
				}
			})
		}
}
