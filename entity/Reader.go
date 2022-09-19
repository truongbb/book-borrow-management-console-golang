package entity

import (
	"bufio"
	"fmt"
	"library-management/constant"
	"os"
	"strings"
)

var countIdReader = 10000

type Reader struct {
	Id         int
	Name       string
	Address    string
	Phone      string
	ReaderType constant.ReaderType
}

func (reader *Reader) InputInfoReader() {
	reader.Id = countIdReader
	stdin := bufio.NewReader(os.Stdin)
	fmt.Printf("Nhập tên bạn đọc: ")
	reader.Name, _ = stdin.ReadString('\n')
	reader.Name = strings.TrimSuffix(reader.Name, "\n")
	fmt.Printf("Nhập tên địa chỉ: ")
	reader.Address, _ = stdin.ReadString('\n')
	reader.Address = strings.TrimSuffix(reader.Address, "\n")
	fmt.Printf("Nhập số điện thoại: ")
	reader.Phone, _ = stdin.ReadString('\n')
	reader.Phone = strings.TrimSuffix(reader.Phone, "\n")
	fmt.Println("Nhập loại bạn đọc: ")
	fmt.Println("1. Sinh viên")
	fmt.Println("2. Học viên cao học")
	fmt.Println("3. Giảng viên")
	fmt.Printf("Xin mời lựa chọn: ")
	var choice int
	for {
		fmt.Scanf("%d", &choice)
		if choice >= 1 || choice <= 3 {
			break
		}
		fmt.Printf("Lựa chọn không hợp lệ, vui lòng chọn lại: ")
	}
	readerTypeInstance := new(constant.ReaderType)
	reader.ReaderType = readerTypeInstance.GetReaderType(choice)
	countIdReader++
}

func (reader Reader) ShowInfoReader() string {
	return fmt.Sprintf("Reader: {Id='%d', Name='%s', Address='%s', Phone='%s', ReaderType='%s'}\n", reader.Id, reader.Name, reader.Address, reader.Phone, reader.ReaderType)
}
