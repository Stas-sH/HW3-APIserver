package server

type Administrator struct {
	Id        int    `json:"admin-id"`
	AdminName string `json:"administrator"`
	Password  string `json:"password"`
}

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"name"`
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func NewAdmin() Administrator {
	return Administrator{}
}

func NewUser() User {
	return User{}
}
