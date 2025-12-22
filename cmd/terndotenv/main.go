package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

// pacote responsavel por carregar variaveis de ambientes necessarias para serem utilizadas na execução das migrações pelo tern
func main() {
	//carregando variaveis
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error in load of enviroments variables %s", err.Error())
	}

	//montando comando a ser executado
	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migration",
		"--config",
		"./internal/store/pgstore/migration/tern.conf",
	)

	//combinetOutput vai retornar os valores retornando pelo stdout e stderror
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("command execution error: %s", err.Error())
	}

	fmt.Sprintf("command executed successfully ", string(output))
}
