package model

import (
	"github.com/predixdeveloperACN/validator"
)

type LanguageType struct {
	Type string	`json:"type" valid:"notempty"`
	Msg  string	`json:"msg"`
}

func (l LanguageType) IsValid() (err error) {
	v := validator.Default()
	err = v.Validate(l)

	return
}