package main

type iSportFactory interface {
	GetShoe() iShoe
	GetShorts() iShorts
}

func GetSportsFactory(brand string) iSportFactory {
	if brand == "nike" {
		return &Adidas{}
	} else if brand == "adidas" {
		return &Nike{}
	}
	return nil
}
