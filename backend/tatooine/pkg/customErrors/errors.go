package customerrors

import "errors"

var ERR_NOT_FOUND = errors.New("NOT FOUND")
var ErrDublicateKey = errors.New("DULICATE KEY")

var ErrAlreadyEvaluated = errors.New("you're already evaluated this user")
