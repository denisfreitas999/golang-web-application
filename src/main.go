package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Produto struct {
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

// Busca e renderiza todos os templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

func connectDB() *sql.DB {
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

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close() // Fechamento da conexão com o banco de dados

	rowsGetProdutos, err := db.Query("SELECT * FROM produtos;")
	if err != nil {
		log.Printf("Erro ao buscar produtos: %v", err)
		http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
		return
	}
	defer rowsGetProdutos.Close() // Fechamento do rowsGetProdutos após a leitura

	var produtos []Produto

	for rowsGetProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = rowsGetProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			log.Printf("Erro ao ler os dados do produto: %v", err)
			http.Error(w, "Erro ao ler os dados do produto", http.StatusInternalServerError)
			return
		}

		p := Produto{
			Id:          id,
			Nome:        nome,
			Descricao:   descricao,
			Preco:       preco,
			Quantidade:  quantidade,
		}

		produtos = append(produtos, p)
	}

	if err := rowsGetProdutos.Err(); err != nil {
		log.Printf("Erro ao iterar sobre os resultados: %v", err)
		http.Error(w, "Erro ao iterar sobre os resultados", http.StatusInternalServerError)
		return
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}