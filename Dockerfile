FROM golang:1.18

WORKDIR /app

# Copy the Go module and Go sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of your application source code to the container
COPY . .

# Build the Go application
RUN go build -o main

# Expose the port that your Gin application will listen on
EXPOSE 8080

# Set the environment variable for Gin to run in production mode
ENV GIN_MODE=release

# Run your Go Gin application when the container starts
CMD ["./main"]
