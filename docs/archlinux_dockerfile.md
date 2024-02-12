Este es el archivo Dockerfile utilizado para probar la imagen de **archlinux** en el proyecto

```Bash
# Base image for container
FROM archlinux:latest

# Set label
LABEL maintainer="manuelgb@correo.ugr.es" \
    version="1.1"

# Set the working directory and create the test directory
WORKDIR /app
RUN mkdir test

# Install Go
RUN pacman -Syu --noconfirm go just

# Create user without privileges for the app
RUN useradd -m --no-log-init -u 1001 ppf-tests
USER ppf-tests

# Copy the needed files
COPY go.mod go.sum justfile ./

# Download the dependencies
RUN go mod download && go mod verify

# Set the test working directory
WORKDIR /app/test

# Set entrypoint for testing
ENTRYPOINT ["just", "test"]
```

No elimino `go.mod` y `go.sum` porque no ocupan mucho espacio y era un rompecabezas intentar hacerlo 
sin darle permisos de escritura al directorio `/app` al usuario `ppf-tests`
