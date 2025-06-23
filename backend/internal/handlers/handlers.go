package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	displayed "github.com/arry/WB_project/internal/model/displayed"
	_ "github.com/lib/pq"
)

func TableHandler(w http.ResponseWriter, r *http.Request) {
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
	actors := []displayed.Displayed{}

	for rows.Next() {
		a := displayed.Displayed{}
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

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		d := displayed.Displayed{
			Familyname: r.FormValue("Familyname"),
			Givenname:  r.FormValue("Givenname"),
			Nation:     r.FormValue("Nation"),
			Number:     r.FormValue("Number"),
			Honorar:    r.FormValue("Honorar"),
		}
		displayed.Insert(d)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		http.ServeFile(w, r, "site/form.html")
	}
}

func PrintHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("Familyname"))
	fmt.Println(r.FormValue("Givenname"))
	fmt.Println(r.FormValue("Nation"))
	fmt.Println(r.FormValue("Number"))
	fmt.Println(r.FormValue("Honorar"))
	fmt.Fprintf(w, "Фамилия: %s Имя: %s", r.FormValue("Familyname"), r.FormValue("Givenname"))

}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(r.FormValue("Familyname"))
		fmt.Println(r.FormValue("Givenname"))
		fmt.Println(r.FormValue("Nation"))
		fmt.Println(r.FormValue("Number"))
		fmt.Println(r.FormValue("Honorar"))
		http.Redirect(w, r, "/table", 301)
	} else {
		http.ServeFile(w, r, "site/form.html")
	}
}

func CreateHandler2(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	d := displayed.Displayed{
		Familyname: r.FormValue("Familyname"),
		Givenname:  r.FormValue("Givenname"),
		Nation:     r.FormValue("Nation"),
		Number:     r.FormValue("Number"),
		Honorar:    r.FormValue("Honorar"),
	}
	displayed.Insert(d)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
