package handlers

import (
	"html/template"
	"log"
	"net/http"

	"repository"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Viewing orders")

	oo, err := repository.LoadOrders()

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "../html/error.html", http.StatusInternalServerError)
	}

	t, _ := template.ParseFiles("./templates/orders.html")

	t.Execute(w, oo)
}
