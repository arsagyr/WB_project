package model

import (
	"database/sql"
	"fmt"
	// names "github.com/arry/WB_project/model/names"
)

var lastid int

type Displayed struct { //Сущность, которая будет выводиться в конце
	Id         int
	Familyname string  // Фамилия актёра
	Givenname  string  // Имя актёра
	Nation     string  // Национальность (гражданство)
	Number     int16   // Число фильмов
	Honorar    float32 // Суммарный гонорар
}

func PrintDisplayed() {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT Actors.id, Names.Family, Names.Given, Nations.Name, Number, Honorar FROM Actors 
		JOIN Names ON Actors.Nameid=Names.id
		JOIN Nations ON Actors.Nationid=Nations.id
		`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	actors := []Displayed{}

	for rows.Next() {
		a := Displayed{}
		err := rows.Scan(&a.Id, &a.Familyname, &a.Givenname, &a.Nation, &a.Number, &a.Honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
	}
	for _, a := range actors {
		fmt.Println(a.Id, "|", a.Familyname, "|", a.Givenname, "|", a.Nation, "|", a.Number, "|", a.Honorar)
	}
}
