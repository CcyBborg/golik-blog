package store

// SortDirection - направление сортировки. true - по возрастанию, false - по убыванию
type SortDirection bool

func (sd SortDirection) String() string {
	if sd {
		return "ASC"
	} else {
		return "DESC"
	}
}

var (
	SortAsc  SortDirection = true
	SortDesc SortDirection = false
)

// Sort - список полей, по которым можно сортировать список постов
type Sort struct {
	UpdatedAt *SortDirection
}

// Pagination - параметры пагинации
type Pagination struct {
	Limit  int64
	Offset int64
}

// Opts - опции получения списка постов
type Opts struct {
	Pagination Pagination
	Sort       Sort
}
