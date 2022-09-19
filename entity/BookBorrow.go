package entity

import "fmt"

type BookBorrow struct {
	Reader           Reader
	BookBorrowDetail []BookBorrowDetail
	TotalBorrowBook  int
}

func NewBookBorrow(reader Reader, bookBorrowDetail []BookBorrowDetail) *BookBorrow {
	return &BookBorrow{Reader: reader, BookBorrowDetail: bookBorrowDetail}
}

func (bookBorrow BookBorrow) ShowBorrowInfo() {
	fmt.Printf("Borrow: { %s, Books: [", bookBorrow.Reader.ShowInfoReader())

	borrowDetails := bookBorrow.BookBorrowDetail
	for i := 0; i < len(borrowDetails); i++ {
		fmt.Printf("%s, ", borrowDetails[i].showBorrowDetailInfo())
	}

	fmt.Printf("]}\n")
}

func (bookBorrow *BookBorrow) CalculateTotalBook() {
	var total = 0
	for i := 0; i < len(bookBorrow.BookBorrowDetail); i++ {
		total = total + bookBorrow.BookBorrowDetail[i].Quantity
	}
	bookBorrow.TotalBorrowBook = total
}
