package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
)

// Memory-stored template
var memoryTemplate *template.Template
var loadOnce sync.Once

// LoadHTMLIntoMemory loads the HTML file into a template in memory
func LoadHTMLIntoMemory(filePath string) error {
	var err error
	loadOnce.Do(func() {
		htmlContent, readErr := os.ReadFile(filePath)
		if readErr != nil {
			err = fmt.Errorf("failed to read file: %v", readErr)
			return
		}
		memoryTemplate, err = template.New("index").Parse(string(htmlContent))
	})
	return err
}

// Handler for serving HTML from memory
func serveFromMemory(w http.ResponseWriter, r *http.Request) {
	if memoryTemplate == nil {
		http.Error(w, "Memory template not initialized", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err := memoryTemplate.Execute(w, map[string]string{"Method": "Memory"})
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

// Handler for serving HTML from disk on each request
func serveFromDisk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Load and parse the template from the file on every request
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Failed to load HTML file from disk", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]string{"Method": "Disk"})
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func main() {
	// Load HTML into memory once when the server starts
	err := LoadHTMLIntoMemory("index.html")
	if err != nil {
		log.Fatalf("Could not load HTML file into memory: %v", err)
	}

	http.HandleFunc("/from-memory", serveFromMemory)
	http.HandleFunc("/from-disk", serveFromDisk)

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
