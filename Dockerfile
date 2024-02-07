# Base image
FROM golang:alpine3.18

# Set label
LABEL maintainer="manuelgb@correo.ugr.es" \
    version="1.0"

# Create user without privileges for the app
RUN adduser --disabled-password -u 1001 ppf-tests

# Set the working directory
WORKDIR /app

# Install task runner
RUN go go install github.com/go-task/task/v3/cmd/task@latest

# Copy the needed files
COPY go.mod go.sum taskfile.yaml ./

# Download the dependencies
RUN go mod download

# Set the test working directory
WORKDIR /app/tests

# Set entrypoint
ENTRYPOINT ["task", "test"]