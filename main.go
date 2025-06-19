package main

import (
	"fmt"

	model "github.com/arry/WB_project/model"
)

func main() {
	// a := model.Actor{
	// 	Id:         1,
	// 	Familyname: 1,
	// 	Givenname:  1,
	// 	Nationid:   1,
	// 	Number:     1,
	// 	Honorar:    1,
	// }
	// a.Print()
	a := model.Actor{
		Familyname: "Doe",
		Givenname:  "John",
		Nation:     "Россия",
		Number:     1,
		Honorar:    1,
	}
	model.InsertActor(a)

	a = model.SelectActorByFN("Doe")
	fmt.Print(a.Familyname)
	fmt.Print(" ")
	fmt.Print(a.Givenname)
}
