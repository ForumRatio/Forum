package forum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func LogPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/logpage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func Welcome(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/welcomepage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func Login(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/connexion.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func Profil(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/profilpage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func SavedProfil(w http.ResponseWriter, r *http.Request, pp *User) {
	var user modifyProfil
	db := InitDatabase("test")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	UpdateUser(db, pp.Id, user.Name, user.Pictures)
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}
func Disconnect(w http.ResponseWriter, r *http.Request, pp *User) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func LoadPostProfile(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("test")
	user := SelectPostrByUser(db, pp.Id)
	//fmt.Println(user)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func LoadUser(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("test")
	user := SelectUserById(db, pp.Id)
	//fmt.Println(user)
	userf, _ := json.Marshal(user)
	//fmt.Println(userf)
	w.Write(userf)
}
func Execute() {
	// db := InitDatabase("test")
	// InsertIntoUsers(db, "moi", "lm", "lm", "lm", 0)
	// InsertIntoSubject(db, "name", 1)
	// InsertIntoSubject(db, "lmlm", 2)
	// InsertIntoSubject(db, "jkl", 3)
	// InsertIntoSubject(db, "njk", 2)
	// InsertIntoContent(db, "Le cheval c'est trop génial", 3, 3, 1)
	//fmt.Println(SelectAllFromSubject(db, 2))
	fmt.Println("http://localhost:8080/")
	dataU := User{1, "", "", "", "", 0}
	PtsU := &dataU
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Welcome(rw, r)
	})
	http.HandleFunc("/logpage", func(rw http.ResponseWriter, r *http.Request) {
		LogPage(rw, r)
	})
	http.HandleFunc("/profil", func(rw http.ResponseWriter, r *http.Request) {
		Profil(rw, r)
	})
	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		Login(rw, r)
	})
	http.HandleFunc("/loadUser", func(rw http.ResponseWriter, r *http.Request) {
		LoadUser(rw, r, PtsU)
	})
	http.HandleFunc("/loadPostUser", func(rw http.ResponseWriter, r *http.Request) {
		LoadPostProfile(rw, r, PtsU)
	})
	http.HandleFunc("/savedProfil", func(rw http.ResponseWriter, r *http.Request) {
		SavedProfil(rw, r, PtsU)
	})
	http.HandleFunc("/disconnect", func(rw http.ResponseWriter, r *http.Request) {
		Disconnect(rw, r, PtsU)
	})
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fi := http.FileServer(http.Dir("./asset/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fi))
	fa := http.FileServer(http.Dir("./src/"))
	http.Handle("/src/", http.StripPrefix("/src/", fa))
	http.ListenAndServe(":8080", nil)
}
