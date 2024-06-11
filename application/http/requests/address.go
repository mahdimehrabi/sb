package requests

import "m1-article-service/domain/entity"

type Address struct {
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
	Street  string `json:"street" validate:"required"`
	ZipCode string `json:"zipCode" validate:"required"`
}

type Addresses []Address

func (addresses Addresses) ToEntities() []*entity.Address {
	addrs := make([]*entity.Address, len(addresses))
	for i, addr := range addresses {
		addrs[i] = &entity.Address{
			ZipCode: addr.ZipCode,
			State:   addr.State,
			Country: addr.Country,
			Street:  addr.Street,
			City:    addr.City,
		}
	}
	return addrs
}
