package types

type Student struct {
	Id    int
	name  string `validate:"required"`
	email string `validate:"required"`
	age   int    `validate:"required"`
}
