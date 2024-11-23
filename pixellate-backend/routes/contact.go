package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// ContactForm represents the structure of the contact form
type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// ContactHandler handles the /api/contact route
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Respond with an error if it's not a POST request
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var contact ContactForm
	// Decode the JSON body into the ContactForm struct
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Simulate saving the contact form (e.g., to a database or email)
	fmt.Printf("New contact form submission: %+v\n", contact)

	// Save the contact form details to a file
	err = saveToFile(contact)
	if err != nil {
		http.Error(w, "Failed to save contact form", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Thank you for contacting Pixellate! We will get back to you within 24 hours.")
}

// saveToFile saves the contact form details to a file
func saveToFile(contact ContactForm) error {
	// Create a directory to store the contact form submissions
	dir := "contact_submissions"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	// Create a unique filename based on the current timestamp
	filename := fmt.Sprintf("contact_%s.txt", time.Now().Format("20060102_150405"))
	filepath := filepath.Join(dir, filename)

	// Open the file for writing
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the contact form details to the file
	_, err = file.WriteString(fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s\n", contact.Name, contact.Email, contact.Message))
	if err != nil {
		return err
	}

	return nil
}
