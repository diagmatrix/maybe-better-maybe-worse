# Listar órdenes por defecto
default:
    @just --list

# Análisis sintáctico del proyecto
check:
    @echo "Comprobando la sintaxis del proyecto..."
    gofmt -e ./pkg/ > /dev/null

# Ejecución de los tests
test:
    @echo "Ejecutando los tests..."
    go test -v ./tests/