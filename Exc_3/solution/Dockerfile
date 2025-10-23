# First step: Build the Go application
FROM golang:1.25 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Execute script that complies the Go application
RUN chmod +x scripts/build-application.sh && ./scripts/build-application.sh

# Second step: Runtime
FROM alpine

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/ordersystem .

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["./ordersystem"]
