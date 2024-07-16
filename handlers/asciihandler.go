package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type AsciiArtData struct {
	Text     string
	AsciiArt string
	Banner   string
}

// AsciiArtHandler handles POST requests, process input text,
// generate ASCII art and render an HTML template with result.
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

		var asciiArtBuffer bytes.Buffer
		for _, words := range textLines {
			for i := 0; i < 8; i++ {
				for _, char := range words {
					asciiArtBuffer.WriteString(lines[int(char-' ')*9+1+i] + " ")
				}
				asciiArtBuffer.WriteString("\n")
			}
			asciiArtBuffer.WriteString("\n")
		}

		asciiArt := asciiArtBuffer.String()

		data := AsciiArtData{Text: text, AsciiArt: asciiArt, Banner: banner}
		tmpl := template.Must(template.ParseFiles("templates/ascii-art.html"))
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error 500: Internal server error", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Error 405: Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ReadBanner reads banner file content and returns splited content as a slice of strings.
func readBanner(banner string) ([]string, error) {
	path := "banners/" + banner
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}
