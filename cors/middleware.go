package cors

import "net/http"

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST,Get,OPTIONS,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Header", "Accepts,Content-Type,Content-Length")
		handler.ServeHTTP(w, r)
	})
}
