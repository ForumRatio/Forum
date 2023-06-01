package forum

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func profil(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/profilpage.html")
	if err != nil {
		log.Fatal(err)
	}

	template.Execute(w, r)
}
func Execute() {
	fmt.Println("http://localhost:8080/")

	http.HandleFunc("/profil", func(rw http.ResponseWriter, r *http.Request) {
		profil(rw, r)
	})
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fi := http.FileServer(http.Dir("./asset/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fi))
	fa := http.FileServer(http.Dir("./src/"))
	http.Handle("/src/", http.StripPrefix("/src/", fa))
	http.ListenAndServe(":8080", nil)
}
