package main

import (
	"fmt"
)

type Article struct {
	Title  string
	Author string
}

func (a Article) ToString() string {
	return fmt.Sprintf("The %q article was written by %v.", a.Title, a.Author)
}

func main() {
	fmt.Println("Hello World")

	my_article := Article{
		Title:  "Studying Go",
		Author: "Me",
	}

	fmt.Println(">>", my_article.Title, "..", my_article.Author)
	fmt.Println(my_article.ToString())

  myMethod()

}
