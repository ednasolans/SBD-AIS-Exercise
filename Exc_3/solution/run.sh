#!/bin/sh


# Stop and remove any existing containers
docker rm -f orderservice postgres 2>/dev/null

# Build the Docker image for the order system
docker build -t orderservice .

# Create a Docker network
docker network create orders-net

# Run the PostgreSQL database container
docker run -d --name postgres --env-file ./debug.env --network orders-net -p 5432:5432 -v pgdata:/var/lib/postgresql/18/docker postgres:18

sleep 10

# Run orderservice container
docker run -d --name orderservice --env-file ./debug.env --network orders-net -p 3000:3000 orderservice