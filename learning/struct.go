package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	MyBook := Books{"test", "we3", "jsk", 234566}

	fmt.Println(Books{"Go", "as", "jiaocheng", 456456})
	fmt.Println(MyBook.author)
}
