package domain

type Validator interface {
	Validate() error
}

type LomsDomain struct {
}

func New() *LomsDomain {
	return &LomsDomain{}
}
