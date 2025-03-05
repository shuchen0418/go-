package main

type AnimalSounder interface {
	MakeDna()
}

func MakeDna(animal AnimalSounder) {
	animal.MakeDna()
}

type Man struct {
	Name string
}

type Woman struct {
	Name string
}

func (m *Man) MakeDna() {
	println("男人的DNA")
}

func (w *Woman) MakeDna() {
	println("女人的DNA")
}

/* func main() {
	MakeDna(&Man{})
	MakeDna(&Woman{})
} */
