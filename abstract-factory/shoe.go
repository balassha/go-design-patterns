package main

type Shoe struct {
	Name string
	Size int
}

func (s Shoe) GetName() string {
	return s.Name
}

func (s Shoe) GetSize() int {
	return s.Size
}
