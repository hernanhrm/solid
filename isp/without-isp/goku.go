package main

type Goku struct{}

func NewGoku() *Goku {
	return &Goku{}
}

func (g Goku) Kick() {
	println("Goku kicks")
}

func (g Goku) Punch() {
	println("Goku punches")
}

func (g Goku) Transform() {
	println("Goku transforms into a Super Saiyan")
}
