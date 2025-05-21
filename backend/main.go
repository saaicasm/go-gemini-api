package main

import (
	"fmt"
	"net/http"
	"strings" // Import strings package for ToTitle
	"text/template"

	"github.com/saaicasm/city-odyssey/genai"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html") // Load the template
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	mux.HandleFunc("/animal", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		animal := r.FormValue("animal")
		fmt.Println("Selected Animal:", animal)

		http.Redirect(w, r, "/animal-page?animal="+animal, http.StatusSeeOther)
	})

	mux.HandleFunc("/animal-page", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/animal.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		animal := r.URL.Query().Get("animal")
		if animal == "" {
			http.Error(w, "Animal not specified", http.StatusBadRequest)
			return
		}

		displayAnimalName := strings.ToTitle(animal) // "raccoon" -> "Raccoon"

		story := genai.ExampleGenerativeModel_GenerateContent_textOnly(animal)

		data := struct {
			Story      string
			AnimalName string
		}{
			Story:      story,
			AnimalName: displayAnimalName,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Starting server on port 8000...")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
