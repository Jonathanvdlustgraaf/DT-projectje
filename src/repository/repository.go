package repository

import (
	"errors"
	"log"

	. "types"
)

const (
	CUSTOMERFILE   = "./data/customers.json"
	PIZZAFILE      = "./data/pizzas.json"
	ORDERFILE      = "./data/orders.json"
	INGREDIENTFILE = "./data/ingredients.json"
)

func GetPizza(name string) (Pizza, error) {
	pp, _ := LoadPizzas()
	for _, p := range pp {
		if p.Name == name {
			return p, nil
		}
	}
	return Pizza{}, errors.New("pizza not found")
}

func LoadPizzas() ([]Pizza, error) {
	var pp []Pizza
	err := loadData(PIZZAFILE, &pp)

	return pp, err
}

func LoadIngredients() (map[string][]Ingredient, error) {
	var ii []Ingredient
	err := loadData(INGREDIENTFILE, &ii)

	var ingredientsMap = make(map[string][]Ingredient)
	for _, i := range ii {
		ingredientsMap[i.Type] = append(ingredientsMap[i.Type], i)
	}
	return ingredientsMap, err
}

func LoadCustomers() ([]Customer, error) {
	var cc []Customer
	err := loadData(CUSTOMERFILE, &cc)

	return cc, err
}

func LoadCustomersWithOrders() ([]CustomerWithOrders, error) {
	cco := []CustomerWithOrders{}
	cc, err := LoadCustomers()
	if err != nil {
		return nil, err
	}
	for _, c := range cc {
		co, _ := LoadOrdersForCustomer(c.Email)
		cco = append(cco, CustomerWithOrders{Customer: c, Orders: co})
	}
	return cco, nil
}

func LoadOrders() ([]Order, error) {
	var oo []Order
	err := loadData(ORDERFILE, &oo)

	return oo, err
}

func LoadOrdersForCustomer(email string) ([]Order, error) {
	oo, err := LoadOrders()
	if err != nil {
		return nil, err
	}
	var co []Order
	for _, o := range oo {
		if o.Email == email {
			co = append(co, o)
		}
	}
	return co, nil
}

func SaveCustomer(c Customer) error {
	log.Println("Saving customer: ", c)
	if c.Email != "" {
		return errors.New("customer not defined")
	}
	cc, err := LoadCustomers()
	if err != nil {
		return err
	}
	for _, ec := range cc {
		if ec.Email == c.Email {
			return errors.New("customer already exists")
		}
	}
	cc = append(cc, c)
	return saveData(CUSTOMERFILE, cc)
}

func SaveOrder(o Order) error {
	log.Println("Saving order: ", o)
	if o.No == 0 || o.Email == "" {
		return errors.New("order not defined")
	}
	oo, err := LoadOrders()
	if err != nil {
		return err
	}
	for _, eo := range oo {
		if eo.No == o.No {
			return errors.New("order already exists")
		}
	}
	oo = append(oo, o)
	return saveData(ORDERFILE, oo)
}

func DeleteCustomer(email string) error {
	cc, err := LoadCustomers()
	if err != nil {
		return err
	}
	for i, ec := range cc {
		if ec.Email == email {
			// found ==> delete, save, return
			newcc := append(cc[:i], cc[i+1:]...)
			return saveData(CUSTOMERFILE, newcc)
		}
	}
	return nil
}

func DeleteOrder(no int) error {
	oo, err := LoadOrders()
	if err != nil {
		return err
	}
	for i, eo := range oo {
		if eo.No == no {
			// found ==> delete, save, return
			newoo := append(oo[:i], oo[i+1:]...)
			return saveData(ORDERFILE, newoo)
		}
	}
	return nil
}
