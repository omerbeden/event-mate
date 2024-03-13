package customerrors

import "errors"

var ERR_NOT_FOUND = errors.New("NOT FOUND")
var ErrDublicateKey = errors.New("DULICATE KEY")
var ErrActivityDoesNotHaveParticipants = errors.New("activity does not have participants")
var ErrAlreadyEvaluated = errors.New("you're already evaluated this user")
var ErrGetLogger = errors.New("failed to get logger")
