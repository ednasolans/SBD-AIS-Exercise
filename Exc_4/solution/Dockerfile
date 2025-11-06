# Use Golang image to build the binary
FROM golang:1.25 AS builder

WORKDIR /app

# Copy the rest of the application code
COPY . /app

# Make sure the build script is executable and run it
RUN chmod +x scripts/build-application.sh
RUN ./scripts/build-application.sh

# Second stage: minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/ordersystem .

EXPOSE 3000

ENV POSTGRES_DB=order
ENV POSTGRES_USER=docker
ENV POSTGRES_PASSWORD=docker
ENV POSTGRES_TCP_PORT=5432
ENV DB_HOST=postgres

CMD ["./ordersystem"]
