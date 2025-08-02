package model

type User struct {
	Name    string
	Age     int
	Sex     string
	Address string
	Id      int
}

func (User) TableName() string {
	return "user"
}
