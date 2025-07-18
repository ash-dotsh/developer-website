package main

import (
    "log"
    "net/http"
    "os"
    "ash-webserver/internal/handlers"
    "ash-webserver/internal/database"
    "ash-webserver/internal/middleware"
    
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize database
    db, err := database.Initialize()
    if err != nil {
        log.Fatal("Failed to initialize database:", err)
    }
    defer db.Close()

    // Initialize handlers
    h := handlers.NewHandlers(db)

    // Setup routes
    r := mux.NewRouter()
    
    // Add middleware
    r.Use(middleware.LoggingMiddleware)
    r.Use(middleware.CORSMiddleware)
    
    // Routes
    r.HandleFunc("/", h.HomeHandler).Methods("GET")
    r.HandleFunc("/health", h.HealthHandler).Methods("GET")
    r.HandleFunc("/api/users", h.GetUsersHandler).Methods("GET")
    r.HandleFunc("/api/users", h.CreateUserHandler).Methods("POST")
    
    // Static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
