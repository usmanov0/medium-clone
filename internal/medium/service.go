package medium

import "fmt"

type Medium struct {
}

func NewMediumService() *Medium {
	return &Medium{}
}

func (m Medium) SomeAction() error {
	fmt.Println("asd")
	return nil
}
