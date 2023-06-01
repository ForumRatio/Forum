package forum

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	template.Execute(w, r)
}
func Execute() {
	db := InitDatabase("test")
	//InsertIntoSubject(db, "kzi", 5)
	fmt.Println(SelectAllFromSubject(db, 2))
	fmt.Println("http://localhost:8080/")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		index(rw, r)
	})
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fi := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fi))
	http.ListenAndServe(":8080", nil)
}
