# Start from a GoLang base image
FROM golang:1.16-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o golang-invoices-rest ./cmd/invoiceserver

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/golang-invoices-rest .

# Expose the port the application listens on
EXPOSE 8080

# Set the entry point for the container
CMD ["./golang-invoices-rest"]
