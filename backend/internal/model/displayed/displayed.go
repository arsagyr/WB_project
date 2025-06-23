package Displayed

import (
	"database/sql"
	"fmt"

	// names "github.com/arry/WB_project/model/names"

	"github.com/arry/WB_project/internal/model/actors"
	"github.com/arry/WB_project/internal/model/names"
	Nation "github.com/arry/WB_project/internal/model/nation"
	_ "github.com/lib/pq"
)

var lastid int

type Displayed struct { //Сущность, которая будет выводиться в конце
	Id         int
	Familyname string // Фамилия актёра
	Givenname  string // Имя актёра
	Nation     string // Национальность (гражданство)
	Number     string // Число фильмов
	Honorar    string // Суммарный гонорар
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

func GetDisplayed() []Displayed {
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
	return actors
}

func Insert(d Displayed) {
	n := names.Name{
		Familyname: d.Familyname,
		Givenname:  d.Givenname,
	}
	nameid := names.Insert(n)
	nationid := Nation.GetID(d.Nation)
	a := actors.Actor{
		Nameid:   nameid,
		Nationid: nationid,
		Number:   d.Number,
		Honorar:  d.Honorar,
	}
	actors.Insert(a)
}

func Search(s1 string, s2 string, s3 string) []Displayed {
	if (s1 == "") && (s2 == "") && (s3 == "") {
		panic("Empty input")
	}

	s := `
		SELECT Actors.id, Names.Family, Names.Given, Nations.Name, Number, Honorar FROM Actors 
		JOIN Names ON Actors.Nameid=Names.id
		JOIN Nations ON Actors.Nationid=Nations.id
	`
	if (s1 != "") && (s2 != "") && (s3 != "") {
		s += `WHERE ((Names.Family LIKE $1) AND (Names.Given LIKE $2)) AND (Nations.Name LIKE $3)`
	} else {
		s += `WHERE ((Names.Family LIKE $1) OR (Names.Given LIKE $2)) OR (Nations.Name LIKE $3)`
	}
	s += `LIMIT 3;`

	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(s, s1, s2, s3)
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
		fmt.Println(a.Id, "|", a.Familyname, "|", a.Givenname, "|", a.Nation, "|", a.Number, "|", a.Honorar)
		actors = append(actors, a)
	}
	return actors
}

func DeleteDisplayed() {

}
