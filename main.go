package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Email string
}

func main() {
	// Configuração da conexão com o banco de dados
	db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1)/perfect_pay")
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco de dados MySQL:", err)
	}
	defer db.Close()

	// Consulta ao banco de dados
	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		log.Fatal("Erro ao executar a consulta no banco de dados MySQL:", err)
	}
	defer rows.Close()

	// Iteração sobre as linhas e impressão do ID e nome separados por hífen
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email)
		if err != nil {
			log.Fatal("Erro ao ler dados da linha:", err)
		}
		fmt.Printf("ID: %d - E-mail: %s\n", user.ID, user.Email)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Erro ao ler todas as linhas:", err)
	}
}
