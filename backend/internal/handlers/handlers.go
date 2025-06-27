package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	model "backend/internal/model"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var nameid int
	row := DB.QueryRow(`
	SELECT Nameid FROM Actors 
	WHERE id=$1
	`, id)
	err := row.Scan(&nameid)
	if err != nil {
		log.Println(err)
	}
	_, err = DB.Exec("DELETE FROM Actors WHERE id = $1", id)
	if err != nil {
		log.Println(err)
	}

	_, err = DB.Exec("DELETE FROM Names WHERE id = $1", nameid)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := DB.QueryRow(`
		SELECT Actors.id, Names.Family, Names.Given, Nations.Name, Number, Honorar FROM Actors 
		JOIN Names ON Actors.Nameid=Names.id
		JOIN Nations ON Actors.Nationid=Nations.id
		WHERE Actors.id = $1
	`, id)
	a := model.Actor{}
	err := row.Scan(&a.Id, &a.Familyname, &a.Givenname, &a.Nation, &a.Number, &a.Honorar)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("internal/templates/edit.html")
		tmpl.Execute(w, a)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	id := r.FormValue(("id"))
	familyname := r.FormValue("familyname")
	givenname := r.FormValue("givenname")
	nation := r.FormValue("nation")
	number := r.FormValue("number")
	honorar := r.FormValue("honorar")

	fmt.Println(id, familyname, givenname, nation, number, honorar)

	_, err = DB.Exec(`
	UPDATE Names SET Family= $1, Given=$2 WHERE id=(
	SELECT nameid FROM Actors WHERE id=$3
	)
	`,
		familyname, givenname, id)
	if err != nil {
		log.Println(err)
	}
	_, err = DB.Exec(`
	UPDATE Actors SET Nationid = (
		SELECT id FROM Nations
		WHERE Name LIKE $1
	),
	Number=$2, 
	Honorar=$3 
	WHERE id = $4
	`, nation, number, honorar, id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 301)
}

// функция добавления данных
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		familyname := r.FormValue("familyname")
		givenname := r.FormValue("givenname")
		nation := "%" + r.FormValue("nation") + "%"
		number := r.FormValue("number")
		honorar := r.FormValue("honorar")

		fmt.Println(familyname, givenname, nation, number, honorar)

		_, err = DB.Exec(`
			INSERT INTO Names (Family, Given)
			VALUES ($1, $2);
		`, familyname, givenname)
		if err != nil {
			log.Println(err)
		}
		row := DB.QueryRow(`
		SELECT id FROM Nations WHERE Name LIKE $1;
		`, nation)
		nationid := 0
		err = row.Scan(&nationid)
		if err != nil {
			log.Println(err)
		}
		_, err = DB.Exec(`
			INSERT INTO Actors (nameid, nationid, number, honorar)
 			VALUES  (
			(SELECT COALESCE(MAX(Id), 0) FROM  Names), 
			$1, $2, $3
			);
			`, nationid, number, honorar)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "internal/templates/create.html")
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
	SELECT Actors.id, Names.Family, Names.Given, Nations.Name, Number, Honorar FROM Actors 
	JOIN Names ON Actors.Nameid=Names.id
	JOIN Nations ON Actors.Nationid=Nations.id
	`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	actors := []model.Actor{}

	for rows.Next() {
		a := model.Actor{}
		err := rows.Scan(&a.Id, &a.Familyname, &a.Givenname, &a.Nation, &a.Number, &a.Honorar)
		if err != nil {
			fmt.Println(err)
			continue
		}
		actors = append(actors, a)
	}

	tmpl, _ := template.ParseFiles("internal/templates/index.html")
	tmpl.Execute(w, actors)
}
