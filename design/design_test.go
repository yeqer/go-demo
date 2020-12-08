package design

import (
	"fmt"
	"testing"
)


func TestDesign(t *testing.T)  {
	var builder PersonBuilder = &shortPersonBuilder{}
	var d *director = &director{personBuilder: builder}
	var p *person = d.build("head(animal)")
	fmt.Println("person build successfully and his head is " + p.getHead())
}


