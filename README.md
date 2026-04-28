
# Gopher Forum 

A lightweight, feature-rich forum application built with **Go** and **SQLite**. This project follows a clean, modular architecture (DDD-lite) to ensure scalability and maintainability.

## Features

- **Authentication & Security**:
  - Secure registration and login using **Bcrypt** for password hashing.
  - Session management using **UUID** and HttpOnly cookies.
- **Posts & Content**:
  - Create and view posts with titles and descriptions.
  - Categorize posts (supports multiple categories per post).
  - Filter posts by category, user, or liked posts.
- **Interactions**:
  - Like or dislike posts and comments.
  - Nested comment system for discussions.
- **Design**:
  - Responsive UI built with Vanilla CSS and HTML templates.
  - Dynamic updates using the internal JSON API.

## 🛠 Tech Stack

- **Backend**: Go (Golang)
- **Database**: SQLite3
- **Frontend**: HTML5, Vanilla CSS, Vanilla JavaScript
- **Security**: Bcrypt, UUID v4

## 📂 Architecture

The project is structured into three distinct layers to decouple business logic from data access and the web interface:

1.  **Handler Layer**: Manages HTTP requests/responses, JSON encoding, and template rendering.
2.  **Service Layer**: The "Brain" of the application. Handles validation, business rules, and coordinates between repositories.
3.  **Repository Layer**: Handles raw SQL queries and interacts directly with the SQLite database.

## 🏁 Getting Started

### Prerequisites

- Go 1.25+ installed on your system.
- GCC (required for the SQLite driver).

### Running Locally

1.  **Clone the repository**:
    ```bash
    git clone https://learn.zone01kisumu.ke/git/sjarso/forum.git
    cd forum
    ```

2.  **Run the application**:
    ```bash
    go run cmd/app/main.go
    ```
    *The server will start at `http://localhost:8080`. The database will be automatically initialized on the first run.*

### Running with Docker

1.  **Build and run the container**:
    ```bash
    docker build -t forum-app .
    docker run -p 8080:8080 forum-app
    ```

## 🏗 Database Schema

The database is managed via automatic migrations. Core tables include:
- `users`: User profiles and credentials.
- `posts`: Core forum content.
- `categories`: Available topics.
- `post_categories`: Links posts to their respective categories.
- `comments`: Discussion threads.
- `reactions`: Like/Dislike tracking.
- `sessions`: Active user sessions.

## 🤝 Contributing

1. Create a feature branch (`feat/your-feature`).
2. Commit changes using descriptive messages.
3. Open a Pull Request for review.
