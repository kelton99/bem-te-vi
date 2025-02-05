# Use the official Golang image
FROM golang:1.22 AS builder

# Instalar os certificados raiz
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Atualizar os certificados
RUN update-ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to cache the dependencies
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux COARCH=amd64 go build -o main cmd/server/main.go

# Use a minimal Docker image to run the Go app
FROM scratch

# Copiar os certificados raiz para a nova imagem
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Run the Go binary
ENTRYPOINT ["./main"]


# docker create network api_network
# docker build -f Dockerfile -t abacateiro .
# docker run -p 8080:8080 abacateiro
# docker run -d -p 8080:8080 --network api_network --name abacateiro abacateiro
