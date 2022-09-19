package constant

type BookSpecialization int

const (
	KH_TN BookSpecialization = 1
	VH_NT BookSpecialization = 2
	DT_VT BookSpecialization = 3
	CNTT  BookSpecialization = 4
)

func (bookSpecialization BookSpecialization) String() string {
	// khai báo một mảng các string
	// toán tử ... để đếm số phần tử
	// số phần tử của mảng là (7)
	names := [...]string{
		"Khoa học tự nhiên",
		"Văn học nghệ thuật",
		"Điện tử viễn thông",
		"Công nghệ thông tin"}

	// `day`: là một trong các giá trị của hằng số Weekday.
	// Nếu hằng số là Sunday, thì day có giá trị là 0.
	// Bắt lỗi trong trường hợp `day` nằm ngoài khoảng của Weekday
	if bookSpecialization < KH_TN || bookSpecialization > CNTT {
		return "Unknown"
	}
	// trả về tên của 1 hằng số Weekday từ mảng names bên trên
	return names[bookSpecialization-1]
}

func (bookSpecialization BookSpecialization) GetBookSpecialization(spec int) BookSpecialization {
	switch spec {
	case 1:
		return KH_TN
	case 2:
		return VH_NT
	case 3:
		return DT_VT
	case 4:
		return CNTT
	}
	return -1
}
