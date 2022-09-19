package entity

import (
	"bufio"
	"fmt"
	"library-management/constant"
	"os"
	"strconv"
	"strings"
)

var countIdBook = 10000

type Book struct {
	Id             int
	Name           string
	Author         string
	Specialization constant.BookSpecialization
	PublishedYear  int
}

func (book *Book) InputInfoBook() {
	book.Id = countIdBook
	stdin := bufio.NewReader(os.Stdin)

	fmt.Printf("Nhập tên sách: ")
	book.Name, _ = stdin.ReadString('\n')
	book.Name = strings.TrimSuffix(book.Name, "\n")

	fmt.Printf("Nhập tác giả của sách: ")
	book.Author, _ = stdin.ReadString('\n')
	book.Author = strings.TrimSuffix(book.Author, "\n")

	fmt.Printf("Nhập năm xuất bản của sách: ")
	var publishedYear, _ = stdin.ReadString('\n')
	book.PublishedYear, _ = strconv.Atoi(strings.TrimSuffix(publishedYear, "\n"))

	fmt.Println("Nhập chuyên ngành của sách: ")
	fmt.Println("1. Khoa học tự nhiên")
	fmt.Println("2. Văn học - Nghệ thuật")
	fmt.Println("3. Điện tử viễn thông")
	fmt.Println("4. Công nghệ thông tin")
	fmt.Printf("Xin mời lựa chọn: ")
	var choice int
	for {
		fmt.Scanf("%d", &choice)
		if choice >= 1 || choice <= 4 {
			break
		}
		fmt.Printf("Lựa chọn không hợp lệ, vui lòng chọn lại: ")
	}
	bookSpecialization := new(constant.BookSpecialization)
	book.Specialization = bookSpecialization.GetBookSpecialization(choice)
	countIdBook++
}

func (book Book) ShowInfoBook() string {
	return fmt.Sprintf("Reader: {Id='%d', Name='%s', Author='%s', Specialization='%s', PublishedYear='%d'}\n", book.Id, book.Name, book.Author, book.Specialization, book.PublishedYear)
}
