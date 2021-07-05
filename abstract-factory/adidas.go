package main

type Adidas struct{}

func (a Adidas) GetShoe() iShoe {
	return &AdidasShoe{
		Shoe{
			Name: "adidas",
			Size: 10,
		},
	}
}

func (a Adidas) GetShorts() iShorts {
	return &AdidasShorts{
		Shorts{
			Name: "adidas",
			Size: "L",
		},
	}
}
