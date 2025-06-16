# 📚 Go Book API

A RESTful CRUD API for managing books, built with idiomatic Go, PostgreSQL.

---

## 📁 Project Structure

```
go-book-api
├── cmd                         # Application entry point (main.go)
│   └── server                  # Starts the HTTP server
├── config                       # Configuration (DB connection, environment, etc.)
├── internal
|   └── books
|       ├── handler.go          # Business Logic Handler
|       ├── model.go            # Book model definition
|       ├── repository.go       # DB access logic
|       └── service.go          # Application logic
├── transport                   # HTTP handlers and route registration
├── Makefile                     # Development scripts (build, test, run)
└── go.mod / go.sum             # Go modules
```

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/go-book-api.git
cd go-book-api
```

### 2. Install dependencies

```bash
make tidy
```

### 3. Run the application

```bash
make run
```

The server will start at: `http://localhost:5555`

---

## 🔧 Environment Setup

Configure database in a separate `.env` file. Copy the contents from `.env.example` to a new file, and rename it to `.env`.

Alter the values in the newly created `.env` file. For instance:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=pass
DB_NAME=books_api
```

Before executing the api, make sure `books` table is created in the connected database. You can run the following SQL command to setup the database.

```SQL
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    published_date DATE NOT NULL
);
```

---

## 📖 API Endpoints

| Method | Endpoint             | Description       |
| ------ | -------------------- | ----------------- |
| GET    | `/api/v1/books`      | Get all books     |
| GET    | `/api/v1/books/{id}` | Get book by ID    |
| POST   | `/api/v1/books`      | Create a new book |
| PUT    | `/api/v1/books/{id}` | Update a book     |
| DELETE | `/api/v1/books/{id}` | Delete a book     |

---

## 📚 Example Book Payload

```json
{
  "title": "Go Programming",
  "author": "Alan A. A. Donovan",
  "published_date": "2023-12-01"
}
```

> If `published_date` is not provided, it defaults to today.

---

## 🛠️ Makefile Commands

```makefile
make run        # Run the app
make build      # Compile the binary
make clean      # Removes the generated binary
make tidy       # Arranges all imports
```

---

## ✅ Tech Stack

* Go (1.23+)
* PostgreSQL
* Standard Library + `database/sql`
* go-base library (public repo by Unbxd)

---

## 📄 License

MIT © 2025 \[Hemanth Kumar]
