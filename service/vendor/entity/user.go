package entity

// User struct
type User struct {
	Username string `xorm:"pk notnull unique" json:"username"`
	Password string `xorm:"notnull" json:"-"`
	Email    string `xorm:"notnull" json:"email"`
	Phone    string `xorm:"notnull" json:"phone"`
}

// Users : Array type of User
type Users []*User
