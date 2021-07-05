package main

type Nike struct{}

func (a Nike) GetShoe() iShoe {
	return &AdidasShoe{
		Shoe{
			Name: "nike",
			Size: 11,
		},
	}
}

func (a Nike) GetShorts() iShorts {
	return &AdidasShorts{
		Shorts{
			Name: "nike",
			Size: "M",
		},
	}
}
