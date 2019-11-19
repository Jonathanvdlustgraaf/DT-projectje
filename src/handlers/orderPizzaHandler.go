package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"helpers"
	"repository"
	. "types"
)

func OrderPizzaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("Ordering pizza")

		r.ParseForm()
		c := Customer{Email: r.FormValue("email"), Name: r.FormValue("name"), Address: r.FormValue("address"), Phone: r.FormValue("phone")}

		// create the customer
		repository.SaveCustomer(c)

		d := time.Now()
		o := Order{No: helpers.GetOrderNo(), Date: d, DateString: d.Format(time.RFC1123Z), Email: c.Email}

		pp := r.Form["pizza"]
		qq := r.Form["qty"]
		for i := 0; i < len(pp); i++ {
			if pp[i] != "" && qq[i] != "" {
				qty, _ := strconv.Atoi(qq[i])
				o.Lines = append(o.Lines, OrderLine{Pizza: pp[i], Qty: qty})
			}
		}
		o.Total = helpers.CalculateTotal(o)

		// create the order
		repository.SaveOrder(o)

		log.Printf("Placing order %v for customer: %v\n ", o, c)

		http.Redirect(w, r, "../html/thankyou.html", http.StatusSeeOther)
	} else {
		log.Println("Preparing pizza order page")

		pp, _ := repository.LoadPizzas()
		// prepare the page
		t, err := template.ParseFiles("./templates/orderpizza.html")
		if err != nil {
			log.Println(err)
		}

		obj := struct {
			Rows   []bool
			Pizzas []Pizza
		}{
			[]bool{true, true, true, true},
			pp,
		}
		t.Execute(w, obj)
	}
}
