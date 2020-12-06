package main

import (
	"fmt"
	"strings"
)

//simple factory design pattern
type Person interface {
	announce()
}

type Man struct {
}

func (Man) announce() {
	fmt.Println("I am a man")
}

type Woman struct {
}

func (Woman) announce() {
	fmt.Println("I am a woman")
}

type Robot struct {
}

func (Robot) announce() {
	fmt.Println("I am a robot")
}

type NvwaFactory struct {
}

func (*NvwaFactory)CreatePerson(like string) Person {
	if strings.EqualFold(like,"M"){
		fmt.Println("NvwaFactory create a Man")
		return &Man{}
	} else if strings.EqualFold(like,"W"){
		fmt.Println("NvwaFactory create a woman")
		return &Woman{}
	} else if strings.EqualFold(like, "R"){
		fmt.Println("NvwaFactory create a Robot")
		return &Robot{}
	}
	fmt.Println("NvwaFactory can't create such a person")
	return nil
}

func main() {
	var nvwaFactory NvwaFactory
	var brand string
	fmt.Scanln(&brand)
	person := nvwaFactory.CreatePerson(brand)
	person.announce()
}
