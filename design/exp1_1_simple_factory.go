package design

import (
	"fmt"
	"strings"
)

//simple factory design pattern

//person parent class
type Person interface {
	announce()
}

//man,woman,robot inherit the parent class separately
type man struct {
}
func (man) announce() {
	fmt.Println("I am a man")
}

type woman struct {
}
func (woman) announce() {
	fmt.Println("I am a woman")
}

type robot struct {
}
func (robot) announce() {
	fmt.Println("I am a robot")
}

//abstract factory
type NvwaFactory struct {
}
func (*NvwaFactory)CreatePerson(like string) Person {
	if strings.EqualFold(like,"M"){
		fmt.Println("NvwaFactory create a man")
		return &man{}
	} else if strings.EqualFold(like,"W"){
		fmt.Println("NvwaFactory create a woman")
		return &woman{}
	} else if strings.EqualFold(like, "R"){
		fmt.Println("NvwaFactory create a robot")
		return &robot{}
	}
	fmt.Println("NvwaFactory can't create such a person")
	return nil
}

