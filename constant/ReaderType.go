package constant

type ReaderType int

const (
	Student ReaderType = 1
	PostUniversity ReaderType = 2
	Teacher ReaderType = 3
)

func (readerType ReaderType) String() string {
	// khai báo một mảng các string
	// toán tử ... để đếm số phần tử
	// số phần tử của mảng là (7)
	names := [...]string{
		"Student",
		"Post-University",
		"Teacher"}

	// `day`: là một trong các giá trị của hằng số Weekday.
	// Nếu hằng số là Sunday, thì day có giá trị là 0.
	// Bắt lỗi trong trường hợp `day` nằm ngoài khoảng của Weekday
	if readerType < Student || readerType > Teacher {
		return "Unknown"
	}
	// trả về tên của 1 hằng số Weekday từ mảng names bên trên
	return names[readerType - 1]
}

func (readerType ReaderType) GetReaderType(rType int) ReaderType {
	switch rType {
	case 1:
		return Student
	case 2:
		return PostUniversity
	case 3:
		return Teacher
	}
	return -1
}