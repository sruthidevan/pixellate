package main

import (
	"pixellate-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	r := gin.Default()

	// Configure CORS
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Update with your frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Apply the CORS middleware
	r.Use(func(c *gin.Context) {
		corsConfig.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})

	// Routes
	r.GET("/api/services", gin.WrapF(routes.ServicesHandler)) // Services route
	r.POST("/api/contact", gin.WrapF(routes.ContactHandler))  // Contact form route

	// Start the server on all interfaces
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
