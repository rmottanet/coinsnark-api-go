# Start from the latest golang base image
FROM golang:latest

# Maintainer Info
LABEL maintainer="Rodrigo Motta rmotta.net@gmail.com"

# Set the Current Working Directory
WORKDIR /

# Copy the source
COPY . .

# Download go modules
RUN go mod download

# Build the Go api
RUN go build -o main ./api/cmd

# Expose the port based on the environment variable
#ARG PORT=8000
#ENV PORT=${PORT}
#EXPOSE ${PORT}
EXPOSE 8080

# 
CMD ["./main"]
