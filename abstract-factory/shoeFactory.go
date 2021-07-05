package main

type iShoe interface {
	GetName() string
	GetSize() int
}

func GetShoe(brand string) iShoe {
	if brand == "nike" {
		return NikeShoe{
			Shoe{Name: brand,
				Size: 10,
			},
		}
	} else if brand == "adidas" {
		return AdidasShoe{
			Shoe{
				Name: brand,
				Size: 11,
			},
		}
	}
	return nil
}
