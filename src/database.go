package forum

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dbname string) *sql.DB {
	db, err := sql.Open("sqlite3", "./database/"+dbname+".db")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	sqlStmt := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER,
		name varchar(255),
		cellphone varchar(255),
		email varchar(255),
		password varchar(255),
		PRIMARY KEY ('id'),
		UNIQUE ('email')
		);
		CREATE TABLE IF NOT EXISTS categories(
			id INTEGER PRIMARY KEY,
			category TEXT
		);
		CREATE TABLE IF NOT EXISTS subject(
			id INTEGER PRIMARY KEY,
			subject TEXT,
			category_id INTEGER NOT NULL,
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);
		CREATE TABLE IF NOT EXISTS posts(
			id INTEGER NOT NULL,
			content varchar(255) NOT NULL,
			subject_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			PRIMARY KEY ('id'),
			FOREIGN KEY (subject_id) REFERENCES subject(id)
			FOREIGN KEY (user_id) REFERENCES users(id)
			);
			PRAGMA foreign_keys = ON;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return db
}

func InsertIntoUsers(db *sql.DB, name string, cellphone string, email string, password string) (int64, error) {
	query1 := `INSERT INTO users ('name','cellphone','email','password')
	Values('` + name + `','` + cellphone + `','` + email + `','` + password + `')
	`
	result, err := db.Exec(query1)
	if err != nil {
		log.Printf("%q: %s\n", err, query1)
		return 0, nil
	}
	return result.LastInsertId()
}
func SelectAllFromUsers(db *sql.DB, table string) []User {
	montre := `SELECT * FROM ` + table
	result, err := db.Query(montre)
	if err != nil {
		log.Printf("%q: %s\n", err, montre)
		return nil
	}
	got := []User{}
	for result.Next() {
		var r User
		err = result.Scan(&r.Id, &r.Name, &r.Cellphone, &r.Email, &r.Password)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		got = append(got, r)
	}
	return got
}
func SelectUserById(db *sql.DB, id int) User {
	montre := `SELECT * FROM users WHERE id = ` + strconv.Itoa(id)
	result := db.QueryRow(montre)
	var result2 User
	err := result.Scan(&result2.Id, &result2.Name, &result2.Cellphone, &result2.Email, &result2.Password)
	if err != nil {
		log.Fatalf("Scan: %v", err)
	}
	return result2
}
func SelectUserNameWithPattern(db *sql.DB, pattern string) []User {
	montre := `SELECT * FROM users WHERE name LIKE '%` + pattern + `%'`
	result, err := db.Query(montre)
	if err != nil {
		log.Printf("%q: %s\n", err, montre)
		return nil
	}
	got := []User{}
	for result.Next() {
		var r User
		err = result.Scan(&r.Id, &r.Name, &r.Cellphone, &r.Email, &r.Password)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		got = append(got, r)
	}
	return got
}
func SelectAllFromSubject(db *sql.DB, cat int) []Subject {
	montre := `SELECT subject.id, subject.subject, subject.category_id FROM subject WHERE subject.category_id = ` + strconv.Itoa(cat)
	result, err := db.Query(montre)
	if err != nil {
		log.Printf("%q: %s\n", err, montre)
		return nil
	}
	got := []Subject{}
	for result.Next() {
		var r Subject
		err = result.Scan(&r.Id, &r.Subject, &r.Category_id)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		got = append(got, r)
	}
	return got
}
func SelectSubjectById(db *sql.DB, id int) Subject {
	montre := `SELECT * FROM subject WHERE id = ` + strconv.Itoa(id)
	result := db.QueryRow(montre)
	var result2 Subject
	err := result.Scan(&result2.Id, &result2.Subject, &result2.Category_id)
	if err != nil {
		log.Fatalf("Scan: %v", err)
	}
	return result2
}
func SelectSubjectWithPattern(db *sql.DB, pattern string) []Subject {
	montre := `SELECT * FROM subject WHERE subject LIKE '%` + pattern + `%'`
	result, err := db.Query(montre)
	if err != nil {
		log.Printf("%q: %s\n", err, montre)
		return nil
	}
	got := []Subject{}
	for result.Next() {
		var r Subject
		err = result.Scan(&r.Id, &r.Subject, &r.Category_id)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		got = append(got, r)
	}
	return got
}
func SelectAllFromPosts(db *sql.DB, cat int) []Posts {
	montre := `SELECT posts.id, posts.content, posts.subject_id, posts.user_id FROM posts WHERE posts.subject_id = ` + strconv.Itoa(cat)
	result, err := db.Query(montre)
	if err != nil {
		log.Printf("%q: %s\n", err, montre)
		return nil
	}
	got := []Posts{}
	for result.Next() {
		var r Posts
		err = result.Scan(&r.Id, &r.Content, &r.Subject_id, &r.User_id)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		got = append(got, r)
	}
	return got
}
func SelectPostrById(db *sql.DB, id int) Posts {
	montre := `SELECT * FROM posts WHERE id = ` + strconv.Itoa(id)
	result := db.QueryRow(montre)
	var result2 Posts
	err := result.Scan(&result2.Id, &result2.Content, &result2.Subject_id, &result2.User_id)
	if err != nil {
		log.Fatalf("Scan: %v", err)
	}
	return result2
}
func SelectPostWithPattern(db *sql.DB, pattern string) []Posts {
	montre := `SELECT * FROM posts WHERE content LIKE '%` + pattern + `%'`
	result, err := db.Query(montre)
	if err != nil {
		log.Printf("%q: %s\n", err, montre)
		return nil
	}
	got := []Posts{}
	for result.Next() {
		var r Posts
		err = result.Scan(&r.Id, &r.Content, &r.Subject_id, &r.User_id)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		got = append(got, r)
	}
	return got
}
func InsertIntoContent(db *sql.DB, content string, subject_id, user_id int) (int64, error) {
	query1 := `INSERT INTO posts ('content','subject_id','user_id')
	Values('` + content + `','` + strconv.Itoa(subject_id) + `','` + strconv.Itoa(user_id) + `')
	`
	result, err := db.Exec(query1)
	if err != nil {
		log.Printf("%q: %s\n", err, query1)
		return 0, nil
	}
	return result.LastInsertId()
}
func InsertIntoSubject(db *sql.DB, subject string, category_id int) (int64, error) {
	query1 := `INSERT INTO subject ('subject','category_id')
	Values('` + subject + `','` + strconv.Itoa(category_id) + `')
	`
	result, err := db.Exec(query1)
	if err != nil {
		log.Printf("%q: %s\n", err, query1)
		return 0, nil
	}
	return result.LastInsertId()
}
