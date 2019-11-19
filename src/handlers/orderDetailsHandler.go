package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"repository"
)

func OrderDetailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Viewing order details")

	tmp := r.URL.Query()["no"]

	orderno, err := strconv.Atoi(tmp[0])
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "../html/error.html", http.StatusBadRequest)
	}

	oo, err := repository.LoadOrders()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "../html/error.html", http.StatusInternalServerError)
	}

	for _, o := range oo {
		if o.No == orderno {
			t, _ := template.ParseFiles("./templates/orderdetails.html")

			t.Execute(w, o)
			return
		}
	}

	http.Redirect(w, r, "../html/error.html", http.StatusNotFound)
}
