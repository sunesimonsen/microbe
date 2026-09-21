package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/sunesimonsen/microbe/config"
)

// VersionETag sets the ETag response header to the version in VERSION when
// that file is available. The file is read when the middleware is created so
// the value remains stable for the lifetime of the server.
func VersionETag(next http.Handler) http.Handler {
	etag := strconv.Quote(config.CurrentVersion)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ETag", etag)
		w.Header().Set("Cache-Control", "public, max-age=3600, must-revalidate")

		if etagMatches(r.Header.Get("If-None-Match"), etag) {
			if r.Method == http.MethodGet || r.Method == http.MethodHead {
				w.WriteHeader(http.StatusNotModified)
			} else {
				w.WriteHeader(http.StatusPreconditionFailed)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

func etagMatches(header, etag string) bool {
	for candidate := range strings.SplitSeq(header, ",") {
		candidate = strings.TrimSpace(candidate)
		if candidate == "*" {
			return true
		}

		// If-None-Match uses weak comparison, so W/"version" matches
		// the strong ETag emitted by this middleware as well.
		candidate = strings.TrimPrefix(candidate, "W/")
		if candidate == etag {
			return true
		}
	}

	return false
}
