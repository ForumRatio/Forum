package forum

type User struct {
	Id        int
	Name      string
	Cellphone string
	Email     string
	Password  string
	Picture   int
}
type modifyProfil struct {
	Pictures int
	Name     string
}
type Category struct {
	Id       int
	Category string
}

type Subject struct {
	Id          int
	Subject     string
	Category_id int
}

type Posts struct {
	Id          int
	Content     string
	Subject_id  int
	Category_id int
	User_id     int
}
