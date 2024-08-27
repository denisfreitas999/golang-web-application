package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

// Busca e renderiza todos os templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul gola V", Preco: 39, Quantidade: 5},
		{"Tenis", "Nike Confortável", 89, 3},
		{"Fone", "Headset Pro", 59,2 },
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}