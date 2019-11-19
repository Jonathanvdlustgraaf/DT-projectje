package types

type Customer struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}