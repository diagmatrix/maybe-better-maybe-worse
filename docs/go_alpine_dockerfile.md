Este es el archivo Dockerfile utilizado para probar la imagen de **golang/alpine** en el proyecto

```Bash
# Base image for container
FROM golang:alpine

# Set label
LABEL maintainer="manuelgb@correo.ugr.es" \
    version="1.1"

# Install task runner
RUN apk add just

# Set the working directory
WORKDIR /app

# Create user for tests
RUN adduser -D -u 1001 ppf-test && chown ppf-test /app
USER ppf-test

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
```