package design

//abstrct class
type person struct {
	head, body, arm, leg string
}
func (p *person)setHead(head string)  {
	p.head = head
}
func (p *person)getHead() string {
	return p.head
}

//抽象建造者
type PersonBuilder interface {
	//由于golang中接口不能含有变量，这里需要使用函数(Build())的方式返回一个变量
	//对应Java代码
	/*
	protected Person p = new Person()
	public Person getPerson(){
		return p;
	}
	 */
	Build() *person

	buildHead(head string) PersonBuilder
}

type shortPersonBuilder struct {
	p *person
}
func (s *shortPersonBuilder)buildHead(head string) PersonBuilder {
	if s.p == nil{
		s.p = &person{}
	}
	s.p.setHead(head)
	return s
}
func (s *shortPersonBuilder) Build() *person{
	return s.p
}

type director struct {
	personBuilder PersonBuilder
}

func (d director)build(head string) *person {
	return d.personBuilder.buildHead(head).Build()
}

/*
//test demo of builder_pattern
func TestDesign(t *testing.T)  {
	var builder PersonBuilder = &shortPersonBuilder{}
	var d *director = &director{personBuilder: builder}
	var p *person = d.build("head(animal)")
	fmt.Println("person build successfully and his head is " + p.getHead())
}
*/