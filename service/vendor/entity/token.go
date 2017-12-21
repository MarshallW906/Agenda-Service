package entity

// Token sturct correlated with the DB
type Token struct {
	Username string `xorm:"pk notnull unique"`
	Token    string `xorm:"notnull unique"`
}
