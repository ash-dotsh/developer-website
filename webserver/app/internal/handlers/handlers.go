package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "log"
)

type Handlers struct {
    db *sql.DB
}

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func NewHandlers(db *sql.DB) *Handlers {
    return &Handlers{db: db}
}

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Ash.it.com</title>
        <style>
            body { font-family: Arial, sans-serif; margin: 40px; }
            .container { max-width: 800px; margin: 0 auto; }
            h1 { color: #333; }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Welcome to Ash.it.com</h1>
            <p>Your GoLang web server is running successfully!</p>
            <p><a href="/health">Check Health</a></p>
            <p><a href="/api/users">View Users API</a></p>
        </div>
    </body>
    </html>`
    w.Write([]byte(html))
}

func (h *Handlers) HealthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
        "service": "ash-webserver",
    })
}

func (h *Handlers) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := h.db.Query("SELECT id, name, email FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            log.Printf("Error scanning user: %v", err)
            continue
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (h *Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
        user.Name, user.Email).Scan(&user.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
