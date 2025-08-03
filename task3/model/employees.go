package model

type Employees struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

func (Employees) TableName() string {
	return "employees"
}
