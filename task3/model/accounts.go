package model

type Accounts struct {
	Id      int
	Balance int
}

func (Accounts) TableName() string {
	return "accounts"
}
