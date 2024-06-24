package handlers

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == "" || banner == "" {
			http.Error(w, "Error 400: Bad request", http.StatusBadRequest)
			return
		}

		lines, err := readBanner(banner)
		if err != nil {
			http.Error(w, "Error 500: Internal server error", http.StatusInternalServerError)
			return
		}

		// Split the input text into multiple lines
		textLines := strings.Split(text, "\r\n")

		// Initialize the ascii art output
		asciiArt := ""

		// Process each line of input text
		for _, words := range textLines {
			for i := 0; i < 8; i++ {
				for _, char := range words {
					asciiArt += lines[int(char-' ')*9+1+i] + " "
				}
				asciiArt += "\n"
			}
			asciiArt += "\n"
		}

		tmpl := template.Must(template.ParseFiles("templates/ascii-art.html"))
		err = tmpl.Execute(w, struct {
			Text     string
			AsciiArt string
		}{
			Text:     text,
			AsciiArt: asciiArt,
		})
		if err != nil {
			http.Error(w, "Error 500: Internal server error", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Error 405: Method not allowed", http.StatusMethodNotAllowed)
	}
}

func readBanner(banner string) ([]string, error) {
	path := "banners/" + banner
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}
