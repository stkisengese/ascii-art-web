package handlers

import (
	"net/http"
	"html/template"
)

// IndexHandler renders an HTML template for the root path.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error 500: Internal server error", http.StatusInternalServerError)
	}
}