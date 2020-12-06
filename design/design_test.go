package design

import "testing"

func TestDesign(t *testing.T)  {
	dbFactory := new(mysqlFactory)
	dbFactory.createConnection().announce()
	dbFactory.createStatement().announce()
}