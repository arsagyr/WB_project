package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	model "github.com/arry/WB_project/internal/model"
	_ "github.com/lib/pq"
)

func CreateHandler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		d := model.Actor{
			Familyname: r.FormValue("Familyname"),
			Givenname:  r.FormValue("Givenname"),
			Nation:     r.FormValue("Nation"),
			Number:     r.FormValue("Number"),
			Honorar:    r.FormValue("Honorar"),
		}
		model.Insert(d)

		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "site/form.html")
	}
}

func IndexHandler2(w http.ResponseWriter, r *http.Request) {
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
	tmpl, _ := template.ParseFiles("site/actorstable.html")
	tmpl.Execute(w, actors)

}
