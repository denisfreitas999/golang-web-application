package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// Carrega o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Lê as variáveis de ambiente
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	// Cria a string de conexão
	conexao := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		user, password, dbName, host, port, sslMode)

	// Tenta abrir a conexão com o banco de dados
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Verifica a conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao pingar no banco de dados: %v", err)
	}

	// Exibe mensagem de sucesso
	fmt.Println("Conectado ao Banco de Dados!")

	return db
}