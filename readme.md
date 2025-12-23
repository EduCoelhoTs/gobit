Uso de migrations => tern

1. criamos um executavel para carregar as variaveis de ambiente e entao executar
    1.1. Criar migration
        // tern new migration_name
    1.2. Executar migration
        // go run cmd/terndotenv/main.go

2. geração de código com sqlc
    // sqlc generate -f ./internal/store/pgstore/sqlc.yml

3. Instalando air para live reloading
    // go install github.com/air-verse/air@latest
    // rodando api: air --build.cmd "go build -o ./bin/api.exe ./cmd/api" --build.bin "./bin/api.exe"
    ou configurar .air.toml e rodar com //air
