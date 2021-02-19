package common

import "fmt"

/*
	组合模式是一种`结构型`设计模式，你可以使用它来将对象组合成树状结构，并且能像使用独立对象那样使用它

*/

type File struct {
	Name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}

func (f *File) getName() string {
	return f.Name
}
