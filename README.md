# 🧵 ThreadForge Forum

A lightweight web forum built with **Go**, **SQLite**, and **Docker** that allows users to register, login, create posts, comment, and interact through likes and dislikes.

---

## 🚀 Features

### 👤 Authentication
- User registration (email, username, password)
- Secure login system
- Password hashing using `bcrypt`
- Cookie-based session management

### 💬 Posts & Comments
- Create posts (registered users only)
- Add comments to posts
- View posts and comments (public access)

### 👍 Interactions
- Like / dislike posts
- Like / dislike comments
- View engagement counts publicly

### 🏷️ Categories & Filtering
- Assign categories to posts
- Filter posts by category
- View user-specific posts and liked posts

---

## 🛠️ Tech Stack

- **Backend:** Go (Golang)
- **Database:** SQLite
- **Authentication:** bcrypt + cookies/sessions
- **Containerization:** Docker
- **Frontend:** HTML, CSS (no frameworks)

---

## 📁 Project Structure
orum/
├── cmd/
│   └── server/
│       └── main.go           # Entry point: Initializes DB and starts the server
├── internal/
│   ├── user/                 # Domain: Registration, Login, Profiles
│   │   ├── handler.go        # HTTP logic (parsing forms)
│   │   ├── service.go        # Business logic (Bcrypt, validation)
│   │   └── repository.go     # SQL queries for 'users' table
│   ├── post/                 # Domain: Threads, Categories, Filtering
│   │   ├── handler.go
│   │   ├── service.go        # Category validation and filtering logic
│   │   └── repository.go     # SQL queries for 'posts' & 'post_categories'
│   ├── comment/              # Domain: Discussion replies
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   ├── interaction/          # Domain: Likes and Dislikes
│   │   ├── service.go        # Toggle logic (prevents double-voting)
│   │   └── repository.go
│   ├── session/              # Domain: Security and Cookies
│   │   ├── service.go        # UUID generation and expiry check
│   │   └── repository.go
│   ├── models/               # Shared Data Structs (User, Post, etc.)
│   └── database/             # SQLite Driver setup & Migration runner
├── migrations/               # SQL files (001_initial_schema.sql)
├── web/
│   ├── static/               # CSS, Images, Vanilla JS
│   └── templates/            # HTML (Layouts and Pages)
├── .gitignore                # Ignore forum.db, binaries, and .env
├── Dockerfile                # Build instructions
├── docker-compose.yml        # Orchestration
├── go.mod                    # Dependency management
└── README.md                 # Setup instructions
