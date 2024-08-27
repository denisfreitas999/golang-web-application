package controllers

import (
	"go-web/src/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// Busca e renderiza todos os templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
    produtos, err := models.GetAllProdutos()
    if err != nil {
        log.Printf("Erro ao buscar produtos: %v", err)
        http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
        return
    }

    temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
    temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        nome := r.FormValue("nome")
        descricao := r.FormValue("descricao")
        preco := r.FormValue("preco")
        quantidade := r.FormValue("quantidade")

        precoConvertido, err := strconv.ParseFloat(preco, 64)
        if err != nil {
            log.Println("Erro na conversão do preço:", err)
            http.Error(w, "Erro na conversão do preço", http.StatusBadRequest)
            return
        }

        quantidadeConvertida, err := strconv.Atoi(quantidade)
        if err != nil {
            log.Println("Erro na conversão da quantidade:", err)
            http.Error(w, "Erro na conversão da quantidade", http.StatusBadRequest)
            return
        }

        err = models.NewProduto(nome, descricao, precoConvertido, quantidadeConvertida)
        if err != nil {
            log.Println("Erro ao cadastrar o produto:", err)
            http.Error(w, "Erro ao cadastrar o produto", http.StatusInternalServerError)
            return
        }

        // Redireciona para a página principal após o sucesso
        http.Redirect(w, r, "/", http.StatusMovedPermanently)
    }
}