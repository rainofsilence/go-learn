package dstruct

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

func RunStruct1() {

	var Book1 Books
	//var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "SuMin"
	Book1.subject = "Go 语言基础教程"
	Book1.bookID = 123112

	printBook(Book1)

	printBook2(&Book1)
}

func printBook(books Books) {
	fmt.Println(books.title, books.author, books.subject, books.bookID)
}

func printBook2(book *Books) {
	fmt.Println(book.title, book.author, book.subject, book.bookID)
}
