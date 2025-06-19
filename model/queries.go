package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PrintActors() {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT Actors.id, Names.Family AS "Family name", Names.Given AS "Given name", Nations.Name AS "Nation", Number, Honorar  FROM Actors 
		JOIN Names ON Actors.Nameid=Names.id
		JOIN Nations ON Actors.Nationid=Nations.id
		`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	actors := []Actor{}

	for rows.Next() {
		a := Actor{}
		err := rows.Scan(&a.Id, &a.Familyname, &a.Givenname, &a.Nation, &a.Number, &a.Honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
	}
	for _, a := range actors {
		fmt.Println(a.Id, a.Familyname, a.Givenname, a.Nation, a.Number, a.Honorar)
	}
}

func SelectActorByFN(s string) Actor {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT Actors.id, Names.Family AS "Family name", Names.Given AS "Given name", Nations.Name AS "Nation", Number, Honorar  FROM Actors 
		JOIN Names ON Actors.Nameid=Names.id
		JOIN Nations ON Actors.Nationid=Nations.id
		WHERE Names.Family LIKE $1 `, s)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	a := Actor{}
	for rows.Next() {
		err := rows.Scan(&a.Id, &a.Familyname, &a.Givenname, &a.Nation, &a.Number, &a.Honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return a
}

func InsertActor(a Actor) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	INSERT INTO Names (Id, Family, Given) 
	VALUES  ((SELECT COALESCE(MAX(Id), 0) + 1 FROM  Names), $1, $2);
		 `, a.Familyname, a.Givenname)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
	result, err = db.Exec(`
	INSERT INTO Actors (Id, Nameid, Nationid, Number, Honorar ) 
	VALUES  ((SELECT COALESCE(MAX(Id), 0) + 1 FROM  Actors), (SELECT COALESCE(MAX(Id), 0) + 1 FROM  Names), (SELECT id FROM Nations WHERE Name LIKE $1), $2, $3);
	`, a.Nation, a.Number, a.Honorar)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err != nil {
		panic(err)
	}
	result.RowsAffected()
}
