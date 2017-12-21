package entity

// User struct
type User struct {
	Username string `xorm:"pk notnull unique"`
	Password string `xorm:"notnull"`
	Email    string `xorm:"notnull"`
	Phone    string `xorm:"notnull"`
}

// Users : Array type of User
type Users []*User
