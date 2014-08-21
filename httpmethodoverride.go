// Package httpmethodoverride provides an http.Handler wrapper that
// allows clients to override the provided HTTP method with a
// `_method` query parameter.
package httpmethodoverride

import "net/http"

// Handler returns an http.Handler that looks for a `_method` URL
// query parameter and updates the *http.Request.Method field with its
// value, and calls the original http.Handler.
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		if method := params.Get("_method"); method != "" {
			r.Method = method
		}

		h.ServeHTTP(w, r)
	})
}
