# Base image for container
FROM alpine:latest

# Set label
LABEL maintainer="manuelgb@correo.ugr.es" \
    version="1.1"

# Install task runner
RUN apk add just

# Install go
RUN apk add go

# Set the working directory
WORKDIR /app

# Create user for tests and set permissions
RUN adduser -D ppf-test && chown ppf-test /app
USER ppf-test

# Set environment variables & permissions so that 'go test' doesn't write on /app/test
ENV GOPATH /home/ppf-test/go
ENV GOCACHE /home/ppf-test/.cache/go-build
RUN mkdir -p /home/ppf-test/.cache/go-build && \
    chmod -R 777 /home/ppf-test/.cache/go-build

# Copy the needed files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download && go mod verify

# Remove unused files
RUN rm go.mod go.sum

# Set the test working directory
WORKDIR /app/test

# Set entrypoint for testing
ENTRYPOINT ["just", "test"]
