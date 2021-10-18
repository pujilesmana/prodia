package utils

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var lock = &sync.Mutex{}
var validate *validator.Validate

func GetValidator() *validator.Validate {
	lock.Lock()
	defer lock.Unlock()

	if validate == nil {
		validate = validator.New()
	}

	return validate
}
