package entity

import "errors"

var (
	ErrRecordNotExist  = errors.New("record not exist")
	ErrHttpStatusNotOk = errors.New("http status is not ok")
)
