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
	id       int
	nameid   int     // Фамилия актёра
	nationid int16   // Национальность (гражданство)
	number   int16   // Число фильмов
	honorar  float32 // Суммарный гонорар
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
		err = rows.Scan(&a.id, &a.nameid, &a.nationid, &a.number, &a.honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
		fmt.Println(actors[i].id, actors[i].nameid, actors[i].nationid, actors[i].number, actors[i].honorar)
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
	VALUES  $1, $2, $3, $4, $5);
		 `, a.id, a.nameid, a.nationid, a.number, a.honorar)
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
		err = rows.Scan(&a.id, &a.nameid, &a.nationid, &a.number, &a.honorar)
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
		err = rows.Scan(&a.id, &a.nameid, &a.nationid, &a.number, &a.honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
		fmt.Println(actors[i].id, actors[i].nameid, actors[i].nationid, actors[i].number, actors[i].honorar)
		i++
	}
}
