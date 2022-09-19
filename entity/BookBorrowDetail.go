package entity

import "fmt"

type BookBorrowDetail struct {
	Book     Book
	Quantity int
}

func NewBookBorrowDetail(book Book, quantity int) *BookBorrowDetail {
	return &BookBorrowDetail{Book: book, Quantity: quantity}
}
func (bookBorrowDetail BookBorrowDetail) showBorrowDetailInfo() string {
	return fmt.Sprintf("BorrowDetail: {%s, quantity: %d}", bookBorrowDetail.Book.ShowInfoBook(), bookBorrowDetail.Quantity)
}