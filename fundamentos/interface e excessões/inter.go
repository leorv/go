package main

import "fmt"

func main() {
	a := Article{
		Title:  "Entendendo interface em Go",
		Author: "Tenille Martins",
	}
	Print(a)
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s", a.Title, a.Author)
}

type Stringer interface {
	String() string
}

type Article struct {
	Title  string
	Author string
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
