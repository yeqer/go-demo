package design

import "fmt"

type rice interface {
	announce()
}

type noodle interface {
	announce()
}

type southRice struct {
}

func (southRice) announce() {
	fmt.Println("I am a southRice")
}

type southNoodle struct {
}

func (southNoodle) announce() {
	fmt.Println("I am a southnoodle")
}

type IFactory interface {
	createRice() rice
	createNoodle() noodle
}

type southFactory struct {
}

func (southFactory) createRice() rice {
	return &southRice{}
}
func (southFactory) createNoodle() noodle {
	return &southNoodle{}
}