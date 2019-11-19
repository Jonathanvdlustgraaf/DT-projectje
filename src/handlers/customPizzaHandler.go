package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	. "types"

	"repository"

	"helpers"
)

func CustomPizzaHandler(writer http.ResponseWriter, request *http.Request) {
	ingredientsMap, _ := repository.LoadIngredients()

	if request.Method == "POST" {

		log.Println("Ordering custom pizza")

		request.ParseForm()

		customer := Customer{Email: request.FormValue("email"), Name: request.FormValue("name"), Address: request.FormValue("address"), Phone: request.FormValue("phone")}
		repository.SaveCustomer(customer)

		date := time.Now()
		order := Order{No: helpers.GetOrderNo(), Date: date, DateString: date.Format(time.RFC1123Z), Email: customer.Email}

		bodem := request.FormValue("bodem")
		saus := request.FormValue("saus")
		toppings := request.Form["toppings"]

		var ingredienten []Ingredient
		ingredienten = append(ingredienten, FindIngredient(bodem, ingredientsMap["bodem"]))
		ingredienten = append(ingredienten, FindIngredient(saus, ingredientsMap["saus"]))

		for i := 0; i < len(toppings); i++ {
			if toppings[i] != "" {
				ingredienten = append(ingredienten, FindIngredient(toppings[i], ingredientsMap["topping"]))
			}
		}

		customPizza := CustomPizza{Ingredienten: ingredienten}

		var orderlines []OrderLine
		orderlines = append(orderlines, OrderLine{Qty: 1, CustomPizza: customPizza})
		order.Lines = orderlines
		order.Total = helpers.CalculateTotal(order)

		// create the order
		repository.SaveOrder(order)

		http.Redirect(writer, request, "../html/thankyou.html", http.StatusSeeOther)

	} else {

		log.Println("Viewing custom pizza")

		template, _ := template.ParseFiles("./templates/CustomPizza.html")

		template.Execute(writer, ingredientsMap)

	}
}

func FindIngredient(name string, ingredients []Ingredient) Ingredient {
	for _, ingredient := range ingredients {
		if ingredient.Name == name {
			return ingredient
		}
	}
	panic("ingredient doesnt exist")
}