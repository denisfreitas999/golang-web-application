package models

import (
	"go-web/src/db"
	"log"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProdutos() ([]Produto, error) {
	db := db.ConnectDB()
	defer db.Close() // Fechamento da conexão com o banco de dados

	rowsGetProdutos, err := db.Query("SELECT id, nome, descricao, preco, quantidade FROM produtos;")
	if err != nil {
        return nil, err
    }
	defer rowsGetProdutos.Close() // Fechamento do rowsGetProdutos após a leitura

	var produtos []Produto

	for rowsGetProdutos.Next() {
        var p Produto
        if err := rowsGetProdutos.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade); err != nil {
            return nil, err
        }
        produtos = append(produtos, p)
    }

	if err := rowsGetProdutos.Err(); err != nil {
        return nil, err
    }
	
	return produtos, nil
}

func NewProduto(nome, descricao string, preco float64, quantidade int) error {
	db := db.ConnectDB()
	defer db.Close()

	inserirProduto, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		log.Println("Erro na conversão da quantidade:", err)
		return err
	}

	inserirProduto.Exec(nome, descricao, preco, quantidade)
	return nil
}

func DeleteProduto(id string) error {
	db := db.ConnectDB()
	defer db.Close()

	deletarProduto, err := db.Prepare("delete from produtos where id = $1;")

	if err != nil {
		log.Println("Erro ao deletar o produto", err)
		return err
	}

	deletarProduto.Exec(id)
	return nil
}