# # Use an official Go runtime as a parent image
FROM golang:1.20 

# # Set the working directory to /app
WORKDIR /app

# # Copy the current directory contents into the container at /app
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy


COPY . .

# # Build the Go application
RUN go build -o main ./cmd/main.go

# # Expose port 8080
EXPOSE 8080

# # Command to run the executable
CMD ["./main", "./cmd/migrate/migrate"]






