package entity

type Address struct {
	ID      int64  `json:"id"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Street  string `json:"street"`
	ZipCode string `json:"zipCode"`
}

func NewAddress(city, state, country, street, zipCode string) *Address {
	return &Address{
		City:    city,
		State:   state,
		Country: country,
		Street:  street,
		ZipCode: zipCode,
	}
}
