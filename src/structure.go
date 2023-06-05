package forum

type User struct {
	Id        int
	Name      string
	Cellphone string
	Email     string
	Password  string
	Picture   int
}
type BoolLogin struct {
	Check string
}
type BoolLogin2 struct {
	Check string
	NS    string
}
type ModifyProfil struct {
	Pictures int
	Name     string
}
type CreateS struct {
	Subject     string
	Question    string
	Category_id int
}
type Checkuser struct {
	Username string
	Password string
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
type Like struct {
	Id      int
	Post_id int
	User_id int
}

type Posts struct {
	Id          int
	Content     string
	Like        int
	Dislike     int
	Fuck        int
	Subject_id  int
	Category_id int
	User_id     int
}
