package main

import (
	"github.com/gastrodon/groudon"
)

type AuthBody struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
	Secret   *string `json:"secret"`
}

func (_ AuthBody) Validators() (values map[string]func(interface{}) (bool, error)) {
	values = map[string]func(interface{}) (bool, error){
		"email":    groudon.ValidEmail,
		"password": groudon.OptionalString,
		"secret":   groudon.OptionalString,
	}

	return
}

func (_ AuthBody) Defaults() (values map[string]interface{}) {
	values = map[string]interface{}{}
	return
}
