package main

import (
	"log"
	"net/http"

	"handlers"
)

func main() {
	log.Println("Starting app on port 5000")

	// handlers
	http.HandleFunc("/view/pizzas", handlers.PizzaHandler)
	http.HandleFunc("/view/orders", handlers.OrderHandler)
	http.HandleFunc("/view/orderdetails", handlers.OrderDetailsHandler)
	http.HandleFunc("/view/customers", handlers.CustomerHandler)
	http.HandleFunc("/view/orderpizza", handlers.OrderPizzaHandler)
	http.HandleFunc("/view/custompizza", handlers.CustomPizzaHandler)

	// static handlers to access files located in /html or /images
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	log.Fatal(http.ListenAndServe(":5000", nil))
}
