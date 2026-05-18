# URL Shortener Microservice

A lightweight and high-performance URL shortener microservice written in Go. It allows you to shorten long URLs and redirects users to the original URL when they access the shortened link.

## 🚀 Features

- **In-Memory Storage**: Fast URL mapping using Go's native map structures.
- **REST API Endpoint**: Simple JSON interface to generate short URL codes.
- **Auto-Redirect**: High-speed redirect using standard HTTP 302 redirect.
- **Dependency Injection**: Clean and maintainable architecture separating handler logic from storage logic.

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Framework**: Standard Library (`net/http`)
- **Database**: In-memory (thread-safe operations coming soon)

---

## 🏁 Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.16 or higher)

### Running the Application

1. Clone this repository:
   ```bash
   git clone https://github.com/jhoni-costa/url-shortener.git
   cd url-shortener
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

The server will start running on: `http://localhost:8080`

---

## 🔌 API Endpoints

### 1. Shorten a URL
Creates a shortened version of a given URL.

- **URL**: `/shorten`
- **Method**: `POST`
- **Headers**: `Content-Type: application/json`
- **Request Body**:
  ```json
  {
    "url": "https://github.com/jhoni-costa"
  }
  ```
- **Response**: `201 Created`
  ```json
  {
    "short_url": "http://localhost:8080/aBcdEf"
  }
  ```

#### Example using `curl`:
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com/jhoni-costa"}'
```

---

### 2. Redirect to Original URL
Redirects a short URL code to its original long URL.

- **URL**: `/:code`
- **Method**: `GET`
- **Response**: `302 Found` (Redirects to original destination)

#### Example using browser or `curl`:
```bash
curl -i http://localhost:8080/aBcdEf
```

---

## 📂 Project Structure

```text
├── handlers/
│   └── url.go      # HTTP handlers for shortening & redirecting
├── storage/
│   └── memory.go   # In-memory storage implementation
├── main.go         # Entry point of the microservice
├── go.mod          # Go module file
└── README.md       # Project documentation
```
