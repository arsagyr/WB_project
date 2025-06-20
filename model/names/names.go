package names

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Name struct {
	Id         int
	Familyname string // Фамилия актёра
	Givenname  string // Имя актёра
}

func CheckName(s1 string, s2 string) bool {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT EXISTS (
	SELECT FROM Names
	WHERE  Family = $1, Given = $2
	);
	`, s1, s2)
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

func Insert(s1 string, s2 string) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	INSERT INTO Names (Id, Family, Given)
	VALUES  ((SELECT COALESCE(MAX(Id), 0) + 1 FROM  Names), $1, $2);
		 `, s1, s2)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
}

func DeleteNames(s1 string, s2 string) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	DELETE FROM Names
	WHERE ((Family LIKE $1) AND (Given LIKE $2));
	`, s1, s2)
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
	DELETE FROM Names
	WHERE id = $1;
	`, id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
}

func Select(id int) Name {
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
	name := Name{}
	names := []Name{}
	for rows.Next() {
		n := Name{}
		err = rows.Scan(&n.Familyname, &n.Givenname)
		if err != nil {
			fmt.Println(err)
			continue
		}
		names = append(names, n)

		fmt.Print(names[0].Familyname, names[0].Givenname)

		name = Name{Familyname: names[0].Familyname, Givenname: names[0].Givenname}
	}
	return name
}

func GetIDs(s1 string, s2 string) []int {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT id FROM Names
	WHERE ((Family LIKE $1) AND (Given LIKE $2));
	`, s1, s2)
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
		fmt.Println(id)
		ids = append(ids, id)
	}
	return ids
}

func PrintNames() {
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
	names := []Name{}
	for rows.Next() {
		n := Name{}
		err = rows.Scan(&n.Id, &n.Familyname, &n.Givenname)
		if err != nil {
			fmt.Println(err)
			continue
		}
		names = append(names, n)
		fmt.Println(names[i].Id, "|", names[i].Familyname, "|", names[i].Givenname)
		i++
	}
}
