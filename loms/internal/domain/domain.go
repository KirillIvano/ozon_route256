package domain

type Validator interface {
	Validate() error
}

type Model struct {
}

func New() *Model {
	return &Model{}
}
