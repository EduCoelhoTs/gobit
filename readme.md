Uso de migrations => tern

1. criamos um executavel para carregar as variaveis de ambiente e entao executar
    // go run cmd/terndotenv/main.go

2. geração de código com sqlc
    // sqlc generate --config ./internal/store/pgstore/sqlc.yml