package design

import "testing"

func TestDesign(t *testing.T)  {
	iFactory := new(southFactory)
	iFactory.createRice().announce()
	iFactory.createNoodle().announce()
}