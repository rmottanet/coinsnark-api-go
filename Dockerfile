# Start from the latest golang base image
FROM golang:latest

# Maintainer Info
LABEL maintainer="Rodrigo Motta rmotta.net@gmail.com"

# Set the Current Working Directory
WORKDIR /api

# Copy the source
COPY . .

# Build the Go api
RUN go build -o main ./api/cmd

# Use env
ENV PORT $PORT

# Expose port
EXPOSE $PORT

# Comando para executar sua aplicação quando o container for iniciado
CMD ["./main"]
