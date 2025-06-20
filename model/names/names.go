package names

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Name struct {
	id         int
	familyname string // Фамилия актёра
	givenname  string // Имя актёра
}

func Create(id int, s1 string, s2 string) Name {
	return Name{
		id:         id,
		familyname: s1,
		givenname:  s2,
	}
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

func Insert(n Name) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	INSERT INTO Names (Id, Family, Given)
	VALUES  ((SELECT COALESCE(MAX(Id), 0) + 1 FROM  Names), $1, $2);
		 `, n.familyname, n.givenname)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
}

func DeleteNames(n Name) {
	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec(`
	DELETE FROM Names
	WHERE ((Family LIKE $1) AND (Given LIKE $2));
	`, n.familyname, n.givenname)
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
	DELETE FROM Names
	WHERE id = $1;
	`, id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result.RowsAffected()
}

func SelectByID(id int) Name {
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
		err = rows.Scan(&n.familyname, &n.givenname)
		if err != nil {
			fmt.Println(err)
			continue
		}
		names = append(names, n)

		fmt.Print(names[0].familyname, names[0].givenname)

		name = Name{familyname: names[0].familyname, givenname: names[0].givenname}
	}
	return name
}

func GetIDs(s1 string, s2 string) []int {
	s := ""
	if s1 == "" {
		s = "SELECT id FROM Names WHERE Given LIKE $2;"
	} else if s2 == "" {
		s = "SELECT id FROM Names WHERE Family LIKE $1;"
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
		err = rows.Scan(&n.id, &n.familyname, &n.givenname)
		if err != nil {
			fmt.Println(err)
			continue
		}
		names = append(names, n)
		fmt.Println(names[i].id, "|", names[i].familyname, "|", names[i].givenname)
		i++
	}
}
