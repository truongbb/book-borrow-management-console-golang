package main

import (
	"fmt"
	"library-management/entity"
	"strings"
)

var readers = make([]entity.Reader, 100)
var books = make([]entity.Book, 100)
var borrows = make([]entity.BookBorrow, 100)

func main() {
	menu()
}

func menu() {
	for {
		var choice = functionChoice()
		switch choice {
		case 1:
			inputNewBook()
			break
		case 2:
			showAllBook()
			break
		case 3:
			inputNewReader()
			break
		case 4:
			showAllReader()
			break
		case 5:
			borrowBook()
			break
		case 6:
			showAllBorrow()
			break
		case 7:
			sortBookBorrow()
			break
		case 8:
			searchBookBorrow()
			break
		case 9:
			return
		}
	}
}

func sortBookBorrow() {
	if isEmptyBorrow() {
		fmt.Printf("Danh sách mượn sách đang trống, vui lòng lên danh sách trước khi sắp xếp!")
		return
	}
	fmt.Println("=================== SẮP XẾP DANH SÁCH MƯỢN SÁCH ===================")
	fmt.Println("1. Theo tên bạn đọc.")
	fmt.Println("2. Theo số lượng cuốn sách được mượn (giảm dần).")
	fmt.Println("===============================================")
	fmt.Print("Xin mời nhập lựa chọn: ")
	var choice int
	for {
		fmt.Scanf("%d", &choice)
		if choice >= 1 && choice <= 2 {
			break
		}
		fmt.Printf("Lựa chọn không hợp lệ, vui lòng chọn lại: ")
	}
	switch choice {
	case 1:
		sortByReaderName()
		showAllBorrow()
		break
	case 2:
		calTotalBorrowBook()
		sortByTotalBook()
		showAllBorrow()
		break
	}

}

func sortByTotalBook() {
	for i := 0; i < len(borrows)-1; i++ {
		for j := i + 1; j < len(borrows); j++ {
			if borrows[i].TotalBorrowBook > borrows[j].TotalBorrowBook {
				var temp = borrows[i]
				borrows[i] = borrows[j]
				borrows[j] = temp
			}
		}
	}
}

func sortByReaderName() {
	for i := 0; i < len(borrows)-1; i++ {
		for j := i + 1; j < len(borrows); j++ {
			if strings.Compare(borrows[i].Reader.Name, borrows[j].Reader.Name) > 0 {
				var temp = borrows[i]
				borrows[i] = borrows[j]
				borrows[j] = temp
			}
		}
	}
}

func calTotalBorrowBook() {
	for i := 0; i < len(borrows); i++ {
		borrows[i].CalculateTotalBook()
	}
}

func searchBookBorrow() {
	if isEmptyBorrow() {
		fmt.Printf("Danh sách mượn sách đang trống, vui lòng lên danh sách trước khi tìm kiếm!")
		return
	}
	fmt.Printf("Nhập tên bạn đọc muốn tìm kiếm: ")
	var readerName = ""
	fmt.Scanf("%s", &readerName)
	for i := 0; i < len(borrows); i++ {
		if borrows[i].Reader != (entity.Reader{}) && strings.Contains(strings.ToLower(borrows[i].Reader.Name), strings.ToLower(readerName)) {
			borrows[i].ShowBorrowInfo()
		}
	}
}

func showAllBorrow() {
	for i := 0; i < len(borrows); i++ {
		if borrows[i].Reader != (entity.Reader{}) {
			borrows[i].ShowBorrowInfo()
		}
	}
}

func borrowBook() {
	if isEmptyBooks() || isEmptyReaders() {
		fmt.Printf("Bạn cần có thông tin về danh sách các cuốn sách và danh sách bạn đọc trước khi thực hiện cho mượn sách\n")
		return
	}

	fmt.Printf("Nhập số lượng bạn đọc muốn mượn sách: ")
	var readerNum = 0
	fmt.Scanf("%d", &readerNum)
	for i := 0; i < readerNum; i++ {
		fmt.Printf("Nhập id bạn đọc thứ %d muốn mượn sách: ", i+1)
		var readerId = 0
		var reader entity.Reader
		for true {
			fmt.Scanf("%d", &readerId)
			reader = searchReaderById(readerId)
			if reader != (entity.Reader{}) {
				break
			}
			fmt.Printf("Không có bạn đọc mang mã %d, xin mời nhập lại: ", readerId)
		}
		fmt.Printf("Nhập số lượng sách mà bạn đọc %s muốn mượn: ", reader.Name)
		var bookNum = 0
		for true {
			fmt.Scanf("%d", &bookNum)
			if bookNum <= 5 {
				break
			}
			fmt.Printf("Không được mượn quá 5 đầu sách cho 1 lượt mượn, xin mời nhập lại: ")
		}

		var borrowDetails = make([]entity.BookBorrowDetail, bookNum)
		var k = 0
		for j := 0; j < bookNum; j++ {
			fmt.Printf("Nhập id sách thứ %d mà bạn đọc này muốn mượn: ", j+1)
			var bookId = 0
			var book entity.Book
			for true {
				fmt.Scanf("%d", &bookId)
				book = searchBookById(bookId)
				if book != (entity.Book{}) {
					break
				}
				fmt.Printf("Không có cuốn sách mang mã %d, xin mời nhập lại: ", bookId)
			}

			var quantity = 0
			fmt.Printf("Nhập số lượng cuốn muốn mượn ở đầu sách này: ")
			for true {
				fmt.Scanf("%d", &quantity)
				if quantity <= 3 {
					break
				}
				fmt.Printf("Không được mượn quá 3 cuốn sách trên 1 đầu sách, xin mời nhập lại: ")
			}

			borrowDetail := entity.NewBookBorrowDetail(book, quantity)
			borrowDetails[k] = *borrowDetail
			k++
		}

		bookBorrow := entity.NewBookBorrow(reader, borrowDetails)
		saveBorrow(*bookBorrow)

	}

}

func isEmptyBorrow() bool {
	for i := 0; i < len(borrows); i++ {
		if borrows[i].Reader != (entity.Reader{}) {
			return false
		}
	}
	return true
}

func isEmptyBooks() bool {
	for i := 0; i < len(books); i++ {
		if books[i] != (entity.Book{}) {
			return false
		}
	}
	return true
}

func isEmptyReaders() bool {
	for i := 0; i < len(readers); i++ {
		if readers[i] != (entity.Reader{}) {
			return false
		}
	}
	return true
}

func saveBorrow(borrow entity.BookBorrow) {
	var check = false
	for i := 0; i < len(borrows); i++ {
		if borrows[i].Reader != (entity.Reader{}) && borrows[i].Reader.Id == borrow.Reader.Id {
			check = true

			var oldArr = borrows[i].BookBorrowDetail
			var oldLen = len(borrows[i].BookBorrowDetail)
			var newLen = oldLen + len(borrow.BookBorrowDetail)
			var newBorrowBook = make([]entity.BookBorrowDetail, newLen)
			var newBorrowBookLen = 0

			for k := 0; k < len(borrow.BookBorrowDetail); k++ {
				var checkExist = false
				for j := 0; j < oldLen; j++ {
					if oldArr[j].Book.Id == borrow.BookBorrowDetail[k].Book.Id {
						oldArr[j].Quantity = oldArr[j].Quantity + borrow.BookBorrowDetail[k].Quantity
						checkExist = true
						break
					}
				}
				if !checkExist {
					newBorrowBook[newBorrowBookLen] = borrow.BookBorrowDetail[k]
					newBorrowBookLen++
				}
			}

			borrows[i].BookBorrowDetail = make([]entity.BookBorrowDetail, oldLen+newBorrowBookLen)
			newLen = len(borrows[i].BookBorrowDetail)
			for j := 0; j < oldLen; j++ {
				borrows[i].BookBorrowDetail[j] = oldArr[j]
			}
			var k = 0
			for j := oldLen; j < newLen; j++ {
				borrows[i].BookBorrowDetail[j] = newBorrowBook[k]
				k++
			}
			break
		}
	}
	if check {
		return
	}
	for i := 0; i < len(borrows); i++ {
		if borrows[i].Reader == (entity.Reader{}) {
			borrows[i] = borrow
			break
		}
	}
}

func searchBookById(bookId int) entity.Book {
	for i := 0; i < len(books); i++ {
		if books[i].Id == bookId {
			return books[i]
		}
	}
	return entity.Book{}
}

func searchReaderById(readerId int) entity.Reader {
	for i := 0; i < len(readers); i++ {
		if readers[i].Id == readerId {
			return readers[i]
		}
	}
	return entity.Reader{}
}

func showAllBook() {
	for i := 0; i < len(books); i++ {
		if books[i] != (entity.Book{}) {
			fmt.Printf("%s", books[i].ShowInfoBook())
		}
	}
}

func inputNewBook() {
	fmt.Printf("Nhập số lượng sách mới tại thư viện: ")
	var bookNum = 0
	fmt.Scanf("%d", &bookNum)
	for i := 0; i < bookNum; i++ {
		fmt.Printf("Nhập thông tin cho sách thứ %d\n", i+1)
		book := entity.Book{}
		book.InputInfoBook()
		saveBook(book)
	}
}

func saveBook(book entity.Book) {
	for i := 0; i < len(books); i++ {
		if books[i] == (entity.Book{}) {
			books[i] = book
			break
		}
	}
}

func showAllReader() {
	for i := 0; i < len(readers); i++ {
		if readers[i] != (entity.Reader{}) {
			fmt.Printf("%s", readers[i].ShowInfoReader())
		}
	}
}

func inputNewReader() {
	fmt.Printf("Nhập số lượng bạn đọc đăng ký mới tại thư viện: ")
	var readerNum = 0
	fmt.Scanf("%d", &readerNum)
	for i := 0; i < readerNum; i++ {
		fmt.Printf("Nhập thông tin cho bạn đọc thứ %d\n", i+1)
		reader := entity.Reader{}
		reader.InputInfoReader()
		saveReader(reader)
	}
}

func saveReader(reader entity.Reader) {
	for i := 0; i < len(readers); i++ {
		if readers[i] == (entity.Reader{}) {
			readers[i] = reader
			break
		}
	}
}

func functionChoice() int {
	fmt.Println("=================== PHẦN MỀM QUẢN LÝ THƯ VIỆN ===================")
	fmt.Println("1. Nhập danh sách đầu sách mới.")
	fmt.Println("2. In ra danh sách các sách trong thư viện.")
	fmt.Println("3. Nhập danh sách bạn đọc mới.")
	fmt.Println("4. In ra danh sách các bạn đọc trong thư viện.")
	fmt.Println("5. Lập bảng quản lý mượn sách.")
	fmt.Println("6. In danh sách mượn sách hiện tại.")
	fmt.Println("7. Sắp xếp danh sách mượn sách.")
	fmt.Println("8. Tìm kiếm danh sách mượn sách theo bạn đọc.")
	fmt.Println("9. Thoát.")
	fmt.Println("===============================================")
	fmt.Print("Xin mời nhập lựa chọn: ")
	var choice int
	for {
		fmt.Scanf("%d", &choice)
		if choice >= 1 && choice <= 9 {
			break
		}
		fmt.Printf("Lựa chọn không hợp lệ, vui lòng chọn lại: ")
	}
	return choice
}
