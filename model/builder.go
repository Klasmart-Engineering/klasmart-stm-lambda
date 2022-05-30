package model

type IBuilder interface {
	Build(input interface{}, output interface{}) error
}
