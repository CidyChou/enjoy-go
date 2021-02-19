package main

type I interface {
	Get() int
}

type people struct {
	id int
}

func main() {

}

func (p *people) Get() int {
	return p.id
}
