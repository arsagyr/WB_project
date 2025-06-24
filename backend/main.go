package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/arry/WB_project/internal/handlers"
	displayed "github.com/arry/WB_project/internal/model/displayed"
	"github.com/arry/WB_project/internal/model/names"
	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	http.HandleFunc("/css/STYLES.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "site/css/STYLES.css")
	})
	http.HandleFunc("/media/nofoto.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "site/media/nofoto.jpg")
	})

	http.HandleFunc("/", handlers.IndexHandler2)
	http.HandleFunc("/create", handlers.CreateHandler2)
	// router := httprouter.New()
	// router.GET("/", handlers.ShowForm)
	// router.POST("/submit", handlers.SubmitForm)

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)

	// d := displayed.Displayed{
	// 	Familyname: "Bor",
	// 	Givenname:  "Alex",
	// 	Nation:     " - ",
	// 	Honorar:    "100",
	// 	Number:     "100",
	// }

	displayed.PrintDisplayed()
	names.PrintAll()
}
