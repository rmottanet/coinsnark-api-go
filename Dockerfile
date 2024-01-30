# Start from the latest golang base image
FROM golang:latest

# Maintainer Info
LABEL maintainer="Rodrigo Motta rmotta.net@gmail.com"

# Set the Current Working Directory
WORKDIR /api

# Copy the source
COPY . .

# Download go modules
RUN go mod download

# Build the Go api
RUN go build -o main ./api/cmd

# Set up an environment variable with a default value.
# This will be overridden by the environment variable from the panel if it's set.
ENV PORT=8000

# Expose a default port. The actual PORT environment variable will be used at runtime to start the server.
# Including EXPOSE instruction with a default port to satisfy the system requirements.
EXPOSE 8000

# Comando para executar sua aplicação quando o container for iniciado
CMD ["./main"]
