package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	// Book 1 specifications
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	// Book 2 specifications
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	// Printing book info
	fmt.Printf( "Book 1 title: %s\nBook 2 title: %s\n", Book1.title, Book2.title)
	fmt.Printf( "Book 1 author: %s\nBook 2 author: %s\n", Book1.author, Book2.author)
	fmt.Printf( "Book 1 subject: %s\nBook 2 subject: %s\n", Book1.subject, Book2.subject)
	fmt.Printf( "Book 1 ID: %d\nBook 2 ID: %d\n", Book1.book_id, Book2.book_id)
}
