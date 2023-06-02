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
	check string
}
type modifyProfil struct {
	Pictures int
	Name     string
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
