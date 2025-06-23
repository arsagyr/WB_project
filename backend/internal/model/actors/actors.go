package actors

import (
	"database/sql"
	"fmt"
)

// import (
// 	"database/sql"
// 	"fmt"
// )

type Actor struct {
	Id       int
	Nameid   int    // Фамилия актёра
	Nationid int16  // Национальность (гражданство)
	Number   string // Число фильмов
	Honorar  string // Суммарный гонорар
}

func SelectByID(id int) []Actor {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT Family, Given FROM Actors
	WHERE id=$1;
	`, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	i := 0
	actors := []Actor{}
	for rows.Next() {
		a := Actor{}
		err = rows.Scan(&a.Id, &a.Nameid, &a.Nationid, &a.Number, &a.Honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
		fmt.Println(actors[i].Id, actors[i].Nameid, actors[i].Nationid, actors[i].Number, actors[i].Honorar)
		i++
	}
	return actors
}

func Insert(a Actor) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	INSERT INTO Actors (id, nameid, nationid, number, honorar)
	VALUES  ((SELECT COALESCE(MAX(Id), 0) + 1 FROM  Actors), $1, $2, $3, $4);
		 `, a.Nameid, a.Nationid, a.Number, a.Honorar)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
}

func Delete(id int) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	DELETE FROM Actors
	WHERE id = $1;
	`, id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
}

func SelectAll() []Actor {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT * FROM Actors
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	i := 0
	actors := []Actor{}
	for rows.Next() {
		a := Actor{}
		err = rows.Scan(&a.Id, &a.Nameid, &a.Nationid, &a.Number, &a.Honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
		i++
	}
	return actors
}

func PrintAll() {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT * FROM Actors
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	i := 0
	actors := []Actor{}
	for rows.Next() {
		a := Actor{}
		err = rows.Scan(&a.Id, &a.Nameid, &a.Nationid, &a.Number, &a.Honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
		fmt.Println(actors[i].Id, actors[i].Nameid, actors[i].Nationid, actors[i].Number, actors[i].Honorar)
		i++
	}
}
