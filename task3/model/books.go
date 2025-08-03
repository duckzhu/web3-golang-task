package model

type Books struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func (Books) TableName() string {
	return "books"
}
