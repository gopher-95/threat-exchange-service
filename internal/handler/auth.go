package handler

import "net/http"

// APIKeyMiddleware возвращает middleware, который проверяет наличие и корректность
// API-ключа в заголовке X-API-Key. Если ключ невереный, возвращает 401 Unauthorized.
func APIKeyMiddleware(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("X-API-Key")
			if key != apiKey {
				writeErrorJSON(w, http.StatusUnauthorized, "Unauthoried", "invalid or missing api key")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
