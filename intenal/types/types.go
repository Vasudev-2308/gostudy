package types

type Student struct {
	Id    int
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}

type Teacher struct {
	Id      int
	Name    string `validate:"required"`
	Email   string `validate:"required"`
	Age     int    `validate:"required"`
	Subject string `validate:"required"`
}
