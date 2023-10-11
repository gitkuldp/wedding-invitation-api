# Use an official Go runtime as a parent image
FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/main .

# Expose port 8080 (adjust as needed)
EXPOSE 8080

# Command to run the executable
CMD ["./app/main"]
