package models

type Problem struct {
	ID         int
	Problem_ID int
	Title      string
	Tags       []string
}

type Revise_Problem struct {
	ID         int
	Problem_ID int
}
