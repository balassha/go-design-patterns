package main

type Shorts struct {
	Name string
	Size string
}

func (s Shorts) GetName() string {
	return s.Name
}

func (s Shorts) GetSize() string {
	return s.Size
}
