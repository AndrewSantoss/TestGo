package model

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	CEP      string `json:"cep"`
	Gender   string `json:"gender"`
	Street   string `json:"street"`
	District string `json:"district"`
	City     string `json:"city"`
	State    string `json:"state"`
	Number   int    `json:"number"`
}
