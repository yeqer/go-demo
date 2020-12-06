package design

import "fmt"

type connection interface {
	announce()
}
type statement interface {
	announce()
}

type mysqlConnection struct {
}

func (mysqlConnection) announce() {
	fmt.Println("I am a mysqlConnection")
}

type mysqlStatement struct {
}

func (mysqlStatement) announce() {
	fmt.Println("I am a mysqlStatement")
}

type DBFactory interface {
	createConnection() connection
	createStatement() statement
}

type mysqlFactory struct {
}

func (mysqlFactory) createConnection() connection {
	return &mysqlConnection{}
}
func (mysqlFactory) createStatement() statement {
	return &mysqlStatement{}
}