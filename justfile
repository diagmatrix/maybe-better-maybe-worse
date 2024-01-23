# Justfile para Pre Prod Fixer (ppf)

# Valida sintácticamente los ficheros
check:
    @echo "Validando sintácticamente el proyecto..."
    go vet ./...