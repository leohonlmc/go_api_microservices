# Start from a base image with Go installed
FROM golang:1.21.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies (if using modules)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Start a new stage from scratch for a smaller final image
FROM alpine:latest  

# Set the working directory in the new stage
WORKDIR /root/

# COPY the .env file here in the final stage
COPY --from=builder /app/.env .

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/myapp .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
