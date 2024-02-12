# Base image for container
FROM golang:alpine

# Set label
LABEL maintainer="manuelgb@correo.ugr.es" \
    version="1.1"

# Set the working directory
WORKDIR /app
RUN mkdir test

# Install task runner
RUN apk add just

# Create user without privileges for the app
RUN adduser --disabled-password -u 1001 ppf-tests
USER ppf-tests

# Copy the needed files
COPY go.mod go.sum justfile ./

# Download the dependencies
RUN go mod download && go mod verify

# Set the test working directory
WORKDIR /app/test

# Set entrypoint for testing
ENTRYPOINT ["just", "test"]
