package main

import "net/http"

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler)	http.Handler {
	return session.LoadAndSave(next)
}

// only necessary if you are NOT building a single-page application (like when using React or Vue)
func (app *application) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.Session.Exists(r.Context(), "userID") {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}