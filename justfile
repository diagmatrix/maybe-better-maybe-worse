# Justfile para Pre Prod Fixer (ppf)

# Orden por defecto
default: build

# Valida sintácticamente los ficheros
check:
    @echo "Validando sintácticamente el proyecto..."
    go vet ./...

# Compila el proyecto
build:
    @echo "Creando ejecutable..."
    go build -C cmd -a -v -o pre-prod-fixer.exe