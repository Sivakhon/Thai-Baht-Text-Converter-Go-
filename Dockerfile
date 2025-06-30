# Start from the official Go image
FROM golang:1.23

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum (if any)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Run the app on container start
CMD ["go", "run", "."]
