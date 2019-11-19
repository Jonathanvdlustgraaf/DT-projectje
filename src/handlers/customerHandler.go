package handlers

import (
	"html/template"
	"log"
	"net/http"

	"repository"
)

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Viewing customers")

	cc, _ := repository.LoadCustomersWithOrders()

	t, _ := template.ParseFiles("./templates/customers.html")

	t.Execute(w, cc)
}
