package dto

// 默认分页
const DEFAULT_PAGE_SIZE = 10

// 默认分页大小
const DEFAULT_PAGE_NUMBER = 1

type PageDTO struct {
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	Page       bool   `json:"-"`
	Status     int    `json:"status"`
	UserId     int    `json:"userId"`
	CategoryId int    `json:"categoryId"`
	Keyword    string `json:"keyword"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}

func (pageDTO PageDTO) PageOffset() int {
	return (pageDTO.GetPage() - 1) * pageDTO.GetSize()
}

func (pageDTO PageDTO) GetPage() int {
	if pageDTO.PageNumber <= 0 {
		return DEFAULT_PAGE_NUMBER
	} else {
		return pageDTO.PageNumber
	}
}

func (pageDTO PageDTO) GetSize() int {
	if pageDTO.PageSize == 0 {
		return DEFAULT_PAGE_SIZE
	} else {
		return pageDTO.PageSize
	}
}
