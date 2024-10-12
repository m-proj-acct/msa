# Start from the official Golang base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .
RUN go mod tidy

# Build the application
RUN go build -o user-service .

# Expose the port the app runs on
EXPOSE 8000

# Run the binary
CMD ["./user-service"]

