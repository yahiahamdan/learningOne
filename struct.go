package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"cake": 1.22, "piencake": 3.99},
		tip:   0,
	}
	return b
}
func (b bill) format() string {
	fs := "Bill breadkdown \n"
	var total float64 = 0
	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%v......$%v \n", k+" , ", v)
		total += v
	}
	fs += fmt.Sprintf("total ***** $%0.2f", total)
	return fs
}
