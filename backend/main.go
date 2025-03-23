package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/saaicasm/city-odyssey/genai"
)

// "fmt"
// "log"
// "os"

// "github.com/joho/godotenv"

var selectedAnimal string

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

		// Parse the form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Get the selected animal
		selectedAnimal = r.FormValue("animal")
		fmt.Println("Selected Animal:", selectedAnimal) // Print it in the terminal

		http.Redirect(w, r, "/animal-page", http.StatusSeeOther)

		// Send response back
		// w.WriteHeader(http.StatusOK)
		// fmt.Fprintf(w, "You selected: %s", animal)
	})

	// mux.HandleFunc("/animal-page", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl, err := template.ParseFiles("templates/animal.html")
	// 	if err != nil {
	// 		http.Error(w, "Error loading template", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	data := struct {
	// 		Animal string
	// 	}{
	// 		Animal: selectedAnimal,
	// 	}

	// 	tmpl.Execute(w, data)
	// 	genai.ExampleGenerativeModel_GenerateContent_textOnly(selectedAnimal)

	// })

	mux.HandleFunc("/animal-page", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/animal.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		// Generate the animal story
		story := genai.ExampleGenerativeModel_GenerateContent_textOnly(selectedAnimal)

		data := struct {
			Story string
		}{
			Story: story, // Pass the generated story to the template
		}

		err = tmpl.Execute(w, data) // Execute the template with the data
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
