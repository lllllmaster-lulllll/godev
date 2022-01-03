package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main() {
	//fmt.Println(Books{ "go语言","www.runoob.com","go 语言教程", 6495407})
	//fmt.Println(Books{title: "go语言",author: "www.runoob.com",subject: "go 语言教程", book_id: 6495407})
	//fmt.Println(Books{title: "go语言",author: "www.runoob.com"})
	var book1 Books
	var book2 Books
	book1.title="go语言"
	book1.author="www.runoob.com"
	book1.subject="go 语言教程"
	book1.book_id=6495407

	book2.title="Python 语言"
	book2.author="www.runoob.com"
	book2.subject="Python 语言教程"
	book2.book_id=6495700

	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 subject : %s\n", book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.book_id)
	fmt.Printf("Book 2 title : %s\n", book2.title)
	fmt.Printf("Book 2 author : %s\n", book2.author)
	fmt.Printf("Book 2 subject : %s\n", book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", book2.book_id)
}
