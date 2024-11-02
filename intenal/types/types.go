package types

type Student struct {
	Id    int64
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}

type Teacher struct {
	Id      int64
	Name    string `validate:"required"`
	Email   string `validate:"required"`
	Age     int    `validate:"required"`
	Subject string `validate:"required"`
}
