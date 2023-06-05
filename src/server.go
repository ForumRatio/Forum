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
func CreateSub(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/createsub.html")
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
func Register(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/inscription.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func Categories(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/categorypage.html")
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

//

func SavedProfil(w http.ResponseWriter, r *http.Request, pp *User) {
	var user ModifyProfil
	db := InitDatabase("test")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	UpdateUser(db, pp.Id, user.Name, user.Pictures)
}
func DeletePost(w http.ResponseWriter, r *http.Request, pp *User) {
	var user BoolLogin
	var b BoolLogin
	checklog := false
	db := InitDatabase("test")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	fmt.Println(user)
	user2 := SelectPostrByUser(db, pp.Id)
	for i := 0; i < len(user2); i++ {
		if user.Check == user2[i].Content {
			checklog = true
			DeletePostFromId(db, user2[i].Id)
		}
	}
	if checklog == true {
		http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
		fmt.Println(pp)
		b.Check = "true"
		b1, _ := json.Marshal(b)
		w.Write(b1)
	}
}
func EditPost(w http.ResponseWriter, r *http.Request, pp *User) {
	var user BoolLogin2
	var b BoolLogin
	checklog := false
	db := InitDatabase("test")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	user2 := SelectPostrByUser(db, pp.Id)
	for i := 0; i < len(user2); i++ {
		if user.Check == user2[i].Content {
			checklog = true
			UpdatePost(db, user2[i].Id, user.NS)
		}
	}
	if checklog == true {
		http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
		fmt.Println(pp)
		b.Check = "true"
		b1, _ := json.Marshal(b)
		w.Write(b1)
	}
}
func SavedSub(w http.ResponseWriter, r *http.Request, pp *User) {
	checke := false
	var user CreateS
	var b BoolLogin
	db := InitDatabase("test")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	table := SelectAllFromSubject(db, user.Category_id)
	for c := 0; c < len(table); c++ {
		if user.Subject == table[c].Subject {
			checke = true
		}
	}
	if checke == false {
		fmt.Println(user)
		if pp.Id > 0 {
			InsertIntoSubject(db, user.Subject, user.Category_id)
			table2 := SelectAllFromSubject(db, user.Category_id)
			InsertIntoContent(db, user.Question, 0, 0, 0, table2[len(table2)-1].Id, user.Category_id, pp.Id)
			http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
			b.Check = "true"
			b1, _ := json.Marshal(b)
			w.Write(b1)
		}
	}
}
func CheckUser(w http.ResponseWriter, r *http.Request, pp *User) {
	checklog := false
	var b BoolLogin
	var user Checkuser
	db := InitDatabase("test")
	user2 := SelectAllFromUsers(db, "users")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	for i := 0; i < len(user2); i++ {
		if user.Username == user2[i].Name && user.Password == user2[i].Password {
			checklog = true
			user3 := SelectUserById(db, user2[i].Id)
			pp.Id = user3.Id
			pp.Cellphone = user3.Cellphone
			pp.Name = user3.Name
			pp.Email = user3.Email
			pp.Password = user3.Password
			pp.Picture = user3.Picture
		}
	}
	if checklog == true {
		http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
		b.Check = "true"
		b1, _ := json.Marshal(b)
		w.Write(b1)
	}
}
func Disconnect(w http.ResponseWriter, r *http.Request, pp *User) {
	pp.Id = 0
	pp.Cellphone = ""
	pp.Name = ""
	pp.Email = ""
	pp.Password = ""
	pp.Picture = 0
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func LoadPostProfile(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("test")
	user := SelectPostrByUser(db, pp.Id)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func CreateUser(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("test")
	var b BoolLogin
	check := false
	var use User2
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &use)
	user := SelectAllFromUsers(db, "users")
	fmt.Println(use)
	for i := 0; i < len(user); i++ {
		if use.Email == user[i].Email {
			check = true
		}
	}
	if check == false {
		InsertIntoUsers(db, use.Name, use.Cellphone, use.Email, use.Password, 0)
		http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
		b.Check = "true"
		b1, _ := json.Marshal(b)
		w.Write(b1)
	}
}
func LoadUser(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("test")
	user := SelectUserById(db, pp.Id)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}

//

func Execute() {
	db := InitDatabase("test")
	// InsertIntoUsers(db, "moi", "lm", "lm", "lm", 0)
	// InsertIntoSubject(db, "name", 1)
	// InsertIntoSubject(db, "lmlm", 2)
	// InsertIntoSubject(db, "jkl", 3)
	// InsertIntoSubject(db, "njk", 2)
	// InsertIntoContent(db, "pas cool", 0, 0, 0, 1, 1, 1)
	fmt.Println(SelectAllFromSubject(db, 2))
	fmt.Println("http://localhost:8080/")
	dataU := User{0, "", "", "", "", 0}
	PtsU := &dataU
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Welcome(rw, r)
	})
	http.HandleFunc("/logpage", func(rw http.ResponseWriter, r *http.Request) {
		LogPage(rw, r)
	})
	http.HandleFunc("/createsub", func(rw http.ResponseWriter, r *http.Request) {
		CreateSub(rw, r)
	})
	http.HandleFunc("/profil", func(rw http.ResponseWriter, r *http.Request) {
		Profil(rw, r)
	})
	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		Login(rw, r)
	})
	http.HandleFunc("/register", func(rw http.ResponseWriter, r *http.Request) {
		Register(rw, r)
	})
	http.HandleFunc("/categorypage", func(rw http.ResponseWriter, r *http.Request) {
		Categories(rw, r)
	})
	http.HandleFunc("/loadUser", func(rw http.ResponseWriter, r *http.Request) {
		LoadUser(rw, r, PtsU)
	})
	http.HandleFunc("/createUser", func(rw http.ResponseWriter, r *http.Request) {
		CreateUser(rw, r, PtsU)
	})
	http.HandleFunc("/loadPostUser", func(rw http.ResponseWriter, r *http.Request) {
		LoadPostProfile(rw, r, PtsU)
	})
	http.HandleFunc("/savedProfil", func(rw http.ResponseWriter, r *http.Request) {
		SavedProfil(rw, r, PtsU)
	})
	http.HandleFunc("/deletePost", func(rw http.ResponseWriter, r *http.Request) {
		DeletePost(rw, r, PtsU)
	})
	http.HandleFunc("/editPost", func(rw http.ResponseWriter, r *http.Request) {
		EditPost(rw, r, PtsU)
	})
	http.HandleFunc("/savedSub", func(rw http.ResponseWriter, r *http.Request) {
		SavedSub(rw, r, PtsU)
	})
	http.HandleFunc("/checkUser", func(rw http.ResponseWriter, r *http.Request) {
		CheckUser(rw, r, PtsU)
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
