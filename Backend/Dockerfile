# Use official Go image
FROM golang:1.24-alpine

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your code
COPY . .

# Build your app
RUN go build -o app .

# Expose port 8080
EXPOSE 8080

# Run the app
CMD ["./app"]
