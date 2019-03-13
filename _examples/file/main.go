package main

import (
	"fmt"
	"strings"

	"github.com/Bhinneka/got"
)

// Order struct
type Order struct {
	OrderID string
	Items   []string
}

func main() {
	myOrders := Order{
		OrderID: "00881119977",
		Items:   []string{"Coffee mix", "Tea", "Mendoan", "Gorengan", "Onde-onde"},
	}

	var target strings.Builder

	err := got.Parse("my_template.tmpl", myOrders, &target, nil, got.ParseTemplateFile)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(target.String())
}
