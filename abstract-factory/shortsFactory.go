package main

type iShorts interface {
	GetName() string
	GetSize() string
}

func GetShorts(brand string) iShorts {
	if brand == "nike" {
		return NikeShorts{
			Shorts{Name: brand,
				Size: "L",
			},
		}
	} else if brand == "adidas" {
		return AdidasShorts{
			Shorts{
				Name: brand,
				Size: "M",
			},
		}
	}
	return nil
}
