package main

import (
	"fmt"
)

type Material struct {
	id             string
	name           string
	comment        string
	cant           int
	units          float32
	units_name     string
	price_per_unit float32
}

type Abertura struct {
	id      int
	name    string
	comment string
	items   []Material
	total   float32
}

func (Abertura) actionMaterialInAbertura(action, Abertura, Material) {
	fmt.Println(Abertura.id)
	if action == "add" {
		fmt.Print("add")
	}
	if action == "del" {
		fmt.Print("del")
	}
	if action == "modify" {
		fmt.Print("modify")
	}
}

type List struct {
	name    string
	comment string
	date    string
	to      string
	items   []Abertura
	total   float32
}
