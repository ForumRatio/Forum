package forum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("User")
	store = sessions.NewCookieStore(key)
)

func LogPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/logpage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func ChatPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	template, err := template.ParseFiles("templates/tchat.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func SubjectPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.FormValue("id"))
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	template, err := template.ParseFiles("templates/subjectpage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func CreateSub(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
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
	// session, _ := store.Get(r, "user")
	// auth := session.Values["auth"]
	// if auth == nil || auth == "" {
	// 	fmt.Println(auth)
	// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
	// }
	template, err := template.ParseFiles("templates/categorypage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func Profil(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
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
			InsertIntoContent(db, user.Question, table2[len(table2)-1].Id, user.Category_id, pp.Id)
			http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
			b.Check = "true"
			b1, _ := json.Marshal(b)
			w.Write(b1)
		}
	}
}
func CheckUser(w http.ResponseWriter, r *http.Request, pp *User) {
	session, _ := store.Get(r, "user")
	checklog := false
	var b BoolLogin
	var user Checkuser
	db := InitDatabase("test")
	user2 := SelectAllFromUsers(db, "users")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	for i := 0; i < len(user2); i++ {
		if user.Username == user2[i].Name {
			if CheckPasswordHash(user.Password, user2[i].Password) == true {
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
	}
	if checklog == true {
		b2, _ := json.Marshal(pp)
		session.Values["auth"] = string(b2)
		session.Save(r, w)
		b.Check = "true"
		b1, _ := json.Marshal(b)
		http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
		w.Write(b1)
	}
}
func CheckUser2(w http.ResponseWriter, r *http.Request, pp *User) {
	session, _ := store.Get(r, "user")
	b2, _ := json.Marshal(pp)
	session.Values["auth"] = string(b2)
	session.Save(r, w)
	http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
}
func Disconnect(w http.ResponseWriter, r *http.Request, pp *User) {
	session, _ := store.Get(r, "user")
	session.Values["auth"] = nil
	session.Save(r, w)
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
func CheckCooks(w http.ResponseWriter, r *http.Request, pp *User) {
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	var user User
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		json.Unmarshal([]byte(auth.(string)), &user)
		pp.Id = user.Id
		pp.Cellphone = user.Cellphone
		pp.Name = user.Name
		pp.Email = user.Email
		pp.Password = user.Password
		pp.Picture = user.Picture
		http.Redirect(w, r, "/categorypage", http.StatusSeeOther)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("test")
	var b BoolLogin
	session, _ := store.Get(r, "user")
	check := false
	var use User2
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &use)
	user := SelectAllFromUsers(db, "users")
	fmt.Println(use)
	for i := 0; i < len(user); i++ {
		if use.Email == user[i].Email && use.Name == user[i].Name {
			check = true
		}
	}
	if check == false {
		password, _ := HashPassword(use.Password)
		InsertIntoUsers(db, use.Name, use.Cellphone, use.Email, password, 0)
		user = SelectAllFromUsers(db, "users")
		user3 := SelectUserById(db, len(user))
		pp.Id = user3.Id
		pp.Cellphone = user3.Cellphone
		pp.Name = user3.Name
		pp.Email = user3.Email
		pp.Password = user3.Password
		pp.Picture = user3.Picture
		b2, _ := json.Marshal(pp)
		session.Values["auth"] = string(b2)
		session.Save(r, w)
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
func LoadSubjects(w http.ResponseWriter, r *http.Request) {
	db := InitDatabase("test")
	cat, _ := strconv.Atoi(r.FormValue("id"))
	user := SelectAllFromSubject(db, cat)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}

//

func Execute() {
	// db := InitDatabase("test")
	// InsertIntoUsers(db, "moi", "lm", "lm", "lm", 0)
	// InsertIntoSubject(db, "name", 1)
	// InsertIntoSubject(db, "lmlm", 2)
	// InsertIntoSubject(db, "jkl", 3)
	// InsertIntoSubject(db, "njk", 2)
	// InsertIntoContent(db, "pas cool", 1, 1, 1)
	// fmt.Println(SelectAllFromSubject(db, 2))

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
	http.HandleFunc("/tchat", func(rw http.ResponseWriter, r *http.Request) {
		ChatPage(rw, r)
	})
	http.HandleFunc("/register", func(rw http.ResponseWriter, r *http.Request) {
		Register(rw, r)
	})
	http.HandleFunc("/categorypage", func(rw http.ResponseWriter, r *http.Request) {
		Categories(rw, r)
	})
	http.HandleFunc("/subjects", func(rw http.ResponseWriter, r *http.Request) {
		SubjectPage(rw, r)
	})
	http.HandleFunc("/loadSubjects", func(rw http.ResponseWriter, r *http.Request) {
		LoadSubjects(rw, r)
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
	http.HandleFunc("/Cooks", func(rw http.ResponseWriter, r *http.Request) {
		CheckCooks(rw, r, PtsU)
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
