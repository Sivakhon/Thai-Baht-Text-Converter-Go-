# Thai Baht Text Converter (Go)

This project provides a Go-based tool to convert decimal numbers into Thai text with currency formatting (Baht and Satang).

## 🧠 Features

- Convert decimal (e.g., `1234.56`) to Thai text (e.g., `หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์`)
- Run as a simple CLI tool
- Expose conversion via HTTP API (using Docker)
- Supports test mode for custom logic testing

---

## 🚀 Getting Started

### Prerequisites

- Go 1.23.5 (or use Docker)
- Docker & Docker Compose (optional)

---

## 🧪 Run Locally

### Run Converter Function (CLI)
### Start API

```bash
go run .
```
### Run in Test Mode
```bash
go run . -test
```
## 🐳 Run with Docker
### Build and Start API

```bash
docker compose up --build
```
The API will be available at:
http://localhost:8080 (modify this if your app defines a different endpoint)

## ✅ Example Request
URL: http://localhost:8080/moneydefault, http://localhost:8080/moneygothaibaht
Method: GET
Header: Content-Type: application/json

```
{
  "money": 212121.21
}
```
🔄 Example Response
```
{
  "text": "สองแสนหนึ่งหมื่นสองพันหนึ่งร้อยยี่สิบเอ็ดบาทยี่สิบเอ็ดสตางค์"
}
```

## 🔧 Project Structure
```
.
├── internal/                  # Internal application code
│   ├── handler/               # HTTP handler functions that process requests and send responses
│   ├── model/                 # Struct definitions and domain models (e.g., request structs)
│   ├── pkg/
│   │   └── dto/               # Data Transfer Objects — input formats used between layers
│   ├── router/                # API route definitions and setup (e.g., using Gin)
│   └── services/              # Business logic layer (Thai text conversion lives here)
├── main.go                    # Application entry point (calls router or handler)
├── go.mod / go.sum            # Go module files (dependencies and versions)
├── Dockerfile                 # Docker setup for building and running the Go app
├── docker-compose.yml         # Docker Compose config for running the app and API
└── README.md                  # Project documentation (you're reading it!)
```

## 📝 License
MIT License — use freely for educational and commercial purposes.

## 🙏 Credits
- [shopspring/decimal](https://github.com/shopspring/decimal)
- [artykaikub/go-thaibaht](https://github.com/artykaikub/go-thaibaht)






