package Displayed

import (
	"database/sql"
	"fmt"

	// names "github.com/arry/WB_project/model/names"

	_ "github.com/lib/pq"
)

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
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	nameid := -1
	rows, err := db.Query(`
	SELECT COALESCE(MAX(Id), 0) + 1 FROM  Names
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&nameid)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	result, err := db.Exec(`
	INSERT INTO Names (Id, Family, Given)
	VALUES ($1, $2, $3 );
		 `, nameid, d.Familyname, d.Givenname)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
	result, err = db.Exec(`
	INSERT INTO Actors (id, nameid, nationid, number, honorar)
	VALUES  ((SELECT COALESCE(MAX(Id), 0) + 1 FROM  Actors), (SELECT id FROM Nations WHERE Name LIKE $1), $2, $3, $4);
		 `, nameid, d.Nation, d.Number, d.Honorar)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
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

func DeleteDisplayed(id int) {
	if id == -1 {
		fmt.Println("Wrong input")
	} else {
		connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err := db.Query(`
		SELECT Nameid FROM Actors
		WHERE id=$1;
		`, id)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		var n []int
		for rows.Next() {
			var nn int
			err = rows.Scan(&nn)
			if err != nil {
				fmt.Println(err)
				continue
			}
			n = append(n, nn)
		}
		if len(n) == 0 {

			fmt.Println("Nothing to delete: there is not the actor")
		} else {
			result, err := db.Exec(`
			DELETE FROM Actors
			WHERE id = $1;
			`, id)
			if err != nil {
				panic(err)
			}
			defer db.Close()
			result.RowsAffected()

			result, err = db.Exec(`
			DELETE FROM Names
			WHERE id = $1;
			`, n[0])
			if err != nil {
				panic(err)
			}
			defer db.Close()
			result.RowsAffected()
		}
	}
}
