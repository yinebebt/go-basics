package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type LinkPreview struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Link        string `json:"link"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Fix: Accept the link as a query parameter, extract nodes, and build the preview.
	// Example link preview
	preview := LinkPreview{
		Title:       "Link Preview",
		Description: "A link preview allows users to view content without opening the link, enabling quick reviews before accessing it.",
		ImageURL:    "https://abs.twimg.com/favicons/twitter.3.ico",
		Link:        "https://linkpreview.scallit.com",
	}

	// Parse and execute the template with the preview data
	t, err := template.ParseFiles("link_preview.html")
	if err != nil {
		fmt.Printf("error: +%v", err)
		return
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, preview); err != nil {
		fmt.Printf("error: +%v", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(tpl.Bytes())
}

func main() {
	http.HandleFunc("/preview", handler)
	http.ListenAndServe(":8081", nil)
}
