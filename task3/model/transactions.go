package model

type Transactions struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        int
}

func (Transactions) TableName() string {
	return "transactions"
}
