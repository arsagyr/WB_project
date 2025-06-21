package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Nation struct {
	id   int
	name string
}

func Create(id int, s string) Nation {
	return Nation{
		id:   id,
		name: s,
	}
}

func IsThere(s string) bool {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT EXISTS (
	SELECT FROM Nations
	WHERE  Name = $1);
	`, s)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var b bool
	for rows.Next() {
		err = rows.Scan(&b)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return b
}

func Insert(n Nation) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	INSERT INTO Nations (id, name)
	VALUES  $1, $2);
		 `, n.id, n.name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
}

func DeleteByID(id int) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	DELETE FROM Nations
	WHERE id = $1;
	`, id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
}

func SelectByID(id int) Nation {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT Family, Given FROM Names
	WHERE id=$1;
	`, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	nation := Nation{}
	nations := []Nation{}
	for rows.Next() {
		n := Nation{}
		err = rows.Scan(&n.id, &n.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		nations = append(nations, n)
		nation = Nation{id: nations[0].id, name: nations[0].name}
	}
	return nation
}

func GetIDs(s1 string, s2 string) []int {
	s := ""
	if (s1 == "") || (s2 == "") {
		s = "SELECT id FROM Names WHERE ((Family LIKE $1) OR (Given LIKE $2));"
	} else {
		s = "SELECT id FROM Names WHERE ((Family LIKE $1) AND (Given LIKE $2));"
	}

	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(s, s1, s2)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	ids := []int{}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// fmt.Println(id)
		ids = append(ids, id)
	}
	return ids
}

func SelectAll() []Nation {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT * FROM Nations
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	nations := []Nation{}
	for rows.Next() {
		n := Nation{}
		err = rows.Scan(&n.id, &n.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		nations = append(nations, n)
	}
	return nations
}

func PrintAll() {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT * FROM Names
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	i := 0
	nations := []Nation{}
	for rows.Next() {
		n := Nation{}
		err = rows.Scan(&n.id, &n.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		nations = append(nations, n)
		fmt.Println(nations[i].id, "|", nations[i].name)
		i++
	}
}
