package entity

type Address struct {
	ID      int64
	City    string
	State   string
	Country string
	Street  string
	ZipCode string
	UserID  int64
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
