package main

import (
	"database/sql"
	"log"
	"net/http"

	"forum/internal/auth"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 1. Connect to SQLite
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Create repository → service → handler
	repo := auth.NewRepository(db)
	service := auth.NewService(repo)
	handler := auth.NewHandler(service)

	// 3. Serve static files (CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// 4. Routes

	// Register
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "web/templates/register.html")
			return
		}
		handler.Register(w, r)
	})

	// Login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "web/templates/login.html")
			return
		}
		handler.Login(w, r)
	})

	// Home
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/home.html")
	})

	// 5. Start server
	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
