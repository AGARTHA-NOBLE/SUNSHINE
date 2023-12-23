package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Listening on Port 8080.")
	http.HandleFunc("/api/v1/verify", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		r.Form.Get("username")
		r.Form.Get("password")
		http.Redirect(w, r, "/frontend/webforms/email_verification.html", http.StatusFound)
	})
	http.ListenAndServe(":8080", nil)
}
