package types

type CustomPizza struct {
	Ingredienten []Ingredient
}

type Ingredient struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
