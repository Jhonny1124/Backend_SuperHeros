package main

func main() {

}

type Biography struct {
	FullName string
}
type Powerstats struct {
	Intelligence int
	Strength     int
	Speed        int
	Durability   int
	Power        int
	Combat       int
}
type Images struct {
	Xs string
	Sm string
	Md string
	Lg string
}
type Superhero struct {
	Name       string
	Powerstats Powerstats
	Biography  Biography
	Images     Images
}
