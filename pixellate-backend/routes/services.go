package routes

import (
	"encoding/json"
	"net/http"
)

// Service represents a Pixellate service
type Service struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

// ServicesHandler handles the /api/services route
func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	// Static list of services
	services := []Service{
		{"Custom Software Development", "We transform your unique ideas into functional, scalable applications tailored to your business needs.", "$3000"},
		{"Backend Systems & APIs", "Crafting reliable and efficient systems that ensure seamless operations for your tech infrastructure.", "$2000"},
		{"Cloud Integrations", "Integrate with leading cloud platforms like Azure and AWS to elevate your business capabilities.", "$2500"},
		{"Consulting Services", "Providing expert advice and technical solutions to tackle your toughest tech challenges.", "$1500"},
	}
	// Set JSON response header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}
