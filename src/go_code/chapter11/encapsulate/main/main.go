package main

import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main() {
	p := model.NewPerson("kangnan")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(*p)
	fmt.Println(p.Name, p.GetAge(), p.GetSal())
}
