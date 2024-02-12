# Base image for container
FROM golang:alpine

# Set label
LABEL maintainer="manuelgb@correo.ugr.es" \
    version="1.1"

# Create user without privileges for the app
RUN adduser --disabled-password -u 1001 ppf-tests

# Set the working directory
WORKDIR /app

# Install task runner
RUN apk add just

# Copy the needed files
COPY go.mod go.sum justfile ./

# Download the dependencies
RUN go mod download && go mod verify

# Remove unused files
RUN rm -f go.mod go.sum

# Set the test working directory
WORKDIR /app/test

# Set entrypoint for testing
ENTRYPOINT ["just", "test"]