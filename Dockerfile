# Use an official Go runtime as a parent image
FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o main /app/cmd/main.go

EXPOSE 8080

CMD ["./main"]


