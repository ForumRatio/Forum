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
func ChatPage(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("EXPLOSION")
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	if SelectAllFromUsers(db, "users") == nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		var user User
		json.Unmarshal([]byte(auth.(string)), &user)
		pp.Id = user.Id
		pp.Cellphone = user.Cellphone
		pp.Name = user.Name
		pp.Email = user.Email
		pp.Password = user.Password
		pp.Picture = user.Picture
	}
	template, err := template.ParseFiles("templates/tchat.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func SubjectPage(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("EXPLOSION")
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	if SelectAllFromUsers(db, "users") == nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		var user User
		json.Unmarshal([]byte(auth.(string)), &user)
		pp.Id = user.Id
		pp.Cellphone = user.Cellphone
		pp.Name = user.Name
		pp.Email = user.Email
		pp.Password = user.Password
		pp.Picture = user.Picture
	}
	template, err := template.ParseFiles("templates/subjectpage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func CreateSub(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("EXPLOSION")
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	if SelectAllFromUsers(db, "users") == nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		var user User
		json.Unmarshal([]byte(auth.(string)), &user)
		pp.Id = user.Id
		pp.Cellphone = user.Cellphone
		pp.Name = user.Name
		pp.Email = user.Email
		pp.Password = user.Password
		pp.Picture = user.Picture
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
	template, err := template.ParseFiles("templates/categorypage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}
func Profil(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("EXPLOSION")
	session, _ := store.Get(r, "user")
	auth := session.Values["auth"]
	if SelectAllFromUsers(db, "users") == nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
	if auth == nil || auth == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		var user User
		json.Unmarshal([]byte(auth.(string)), &user)
		pp.Id = user.Id
		pp.Cellphone = user.Cellphone
		pp.Name = user.Name
		pp.Email = user.Email
		pp.Password = user.Password
		pp.Picture = user.Picture
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
	db := InitDatabase("EXPLOSION")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	UpdateUser(db, pp.Id, user.Name, user.Pictures)
}
func SavedPost(w http.ResponseWriter, r *http.Request, pp *User) {
	var user Post2
	db := InitDatabase("EXPLOSION")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	InsertIntoContent(db, user.Check, user.Subject_id, user.Category_id, pp.Id)
}
func Savedlike(w http.ResponseWriter, r *http.Request, pp *User) {
	var user Like2
	db := InitDatabase("EXPLOSION")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	like := SelectLikeById2(db, "like", user.Post_id, user.User_id)
	if like.Id == 0 {
		InsertIntoLike(db, "like", user.Post_id, user.User_id)
	} else {
		DeleteLikeFromId(db, "like", like.Id)
	}
}
func Savedislike(w http.ResponseWriter, r *http.Request, pp *User) {
	var user Like2
	db := InitDatabase("EXPLOSION")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	like := SelectLikeById2(db, "dislike", user.Post_id, user.User_id)
	if like.Id == 0 {
		InsertIntoLike(db, "dislike", user.Post_id, user.User_id)
	} else {
		DeleteLikeFromId(db, "dislike", like.Id)
	}
}
func Savedfuck(w http.ResponseWriter, r *http.Request, pp *User) {
	var user Like2
	db := InitDatabase("EXPLOSION")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	like := SelectLikeById2(db, "fuck", user.Post_id, user.User_id)
	if like.Id == 0 {
		InsertIntoLike(db, "fuck", user.Post_id, user.User_id)
	} else {
		DeleteLikeFromId(db, "fuck", like.Id)
	}
}
func DeletePost(w http.ResponseWriter, r *http.Request, pp *User) {
	var user BoolLogin
	var b BoolLogin
	checklog := false
	db := InitDatabase("EXPLOSION")
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	fmt.Println(user)
	user2 := SelectPostrByUser(db, pp.Id)
	for i := 0; i < len(user2); i++ {
		if user.Check == user2[i].Content {
			checklog = true
			DeleteLikeFromId2(db, "like", user2[i].Id)
			DeleteLikeFromId2(db, "dislike", user2[i].Id)
			DeleteLikeFromId2(db, "fuck", user2[i].Id)
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
	db := InitDatabase("EXPLOSION")
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
	db := InitDatabase("EXPLOSION")
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
	db := InitDatabase("EXPLOSION")
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
	db := InitDatabase("EXPLOSION")
	user := SelectPostrByUser(db, pp.Id)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func LoadPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	db := InitDatabase("EXPLOSION")
	user := SelectAllFromPosts(db, id)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func LoadLike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	db := InitDatabase("EXPLOSION")
	user := SelectLikeById(db, "like", id)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func LoadDislike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	db := InitDatabase("EXPLOSION")
	user := SelectLikeById(db, "dislike", id)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func LoadFuck(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	db := InitDatabase("EXPLOSION")
	user := SelectLikeById(db, "fuck", id)
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
	db := InitDatabase("EXPLOSION")
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
	db := InitDatabase("EXPLOSION")
	user := SelectUserById(db, pp.Id)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func LoadAllUser(w http.ResponseWriter, r *http.Request, pp *User) {
	db := InitDatabase("EXPLOSION")
	user := SelectAllFromUsers(db, "users")
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func LoadSubjects(w http.ResponseWriter, r *http.Request) {
	db := InitDatabase("EXPLOSION")
	cat, _ := strconv.Atoi(r.FormValue("id"))
	user := SelectAllFromSubject(db, cat)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}
func LoadSubjectsTitle(w http.ResponseWriter, r *http.Request) {
	db := InitDatabase("EXPLOSION")
	cat, _ := strconv.Atoi(r.FormValue("id"))
	user := SelectSubjectById(db, cat)
	userf, _ := json.Marshal(user)
	w.Write(userf)
}

//

func Execute() {
	// db := InitDatabase("EXPLOSION")
	// fmt.Println(db)
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
		CreateSub(rw, r, PtsU)
	})
	http.HandleFunc("/profil", func(rw http.ResponseWriter, r *http.Request) {
		Profil(rw, r, PtsU)
	})
	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		Login(rw, r)
	})
	http.HandleFunc("/tchat", func(rw http.ResponseWriter, r *http.Request) {
		ChatPage(rw, r, PtsU)
	})
	http.HandleFunc("/register", func(rw http.ResponseWriter, r *http.Request) {
		Register(rw, r)
	})
	http.HandleFunc("/categorypage", func(rw http.ResponseWriter, r *http.Request) {
		Categories(rw, r)
	})
	http.HandleFunc("/subjects", func(rw http.ResponseWriter, r *http.Request) {
		SubjectPage(rw, r, PtsU)
	})
	http.HandleFunc("/loadSubjects", func(rw http.ResponseWriter, r *http.Request) {
		LoadSubjects(rw, r)
	})
	http.HandleFunc("/loadSubjects2", func(rw http.ResponseWriter, r *http.Request) {
		LoadSubjectsTitle(rw, r)
	})
	http.HandleFunc("/loadUser", func(rw http.ResponseWriter, r *http.Request) {
		LoadUser(rw, r, PtsU)
	})
	http.HandleFunc("/loadAllUser", func(rw http.ResponseWriter, r *http.Request) {
		LoadAllUser(rw, r, PtsU)
	})
	http.HandleFunc("/loadPost", func(rw http.ResponseWriter, r *http.Request) {
		LoadPost(rw, r)
	})
	http.HandleFunc("/loadLike", func(rw http.ResponseWriter, r *http.Request) {
		LoadLike(rw, r)
	})
	http.HandleFunc("/loadDislike", func(rw http.ResponseWriter, r *http.Request) {
		LoadDislike(rw, r)
	})
	http.HandleFunc("/loadFuck", func(rw http.ResponseWriter, r *http.Request) {
		LoadFuck(rw, r)
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
	http.HandleFunc("/savedPost", func(rw http.ResponseWriter, r *http.Request) {
		SavedPost(rw, r, PtsU)
	})
	http.HandleFunc("/savedLike", func(rw http.ResponseWriter, r *http.Request) {
		Savedlike(rw, r, PtsU)
	})
	http.HandleFunc("/savedDislike", func(rw http.ResponseWriter, r *http.Request) {
		Savedislike(rw, r, PtsU)
	})
	http.HandleFunc("/savedFuck", func(rw http.ResponseWriter, r *http.Request) {
		Savedfuck(rw, r, PtsU)
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
