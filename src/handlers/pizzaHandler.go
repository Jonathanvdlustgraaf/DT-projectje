package handlers

import (
	"html/template"
	"log"
	"net/http"

	"repository"
)

func PizzaHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Viewing pizzas")

	pp, _ := repository.LoadPizzas()

	t, _ := template.ParseFiles("./templates/pizzas.html")

	t.Execute(w, pp)
}
