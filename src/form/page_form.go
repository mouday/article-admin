package form

type PageForm struct {
	Page   int `json:"page"`
	Size   int `json:"size"`
	Status int `json:"status"`
	UserId int `json:"userId"`
}

func (pageForm PageForm) PageOffset() int {
	return (pageForm.GetPage() - 1) * pageForm.GetSize()
}

func (pageForm PageForm) GetPage() int {
	if pageForm.Page <= 0 {
		return 1
	} else {
		return pageForm.Page
	}
}

func (pageForm PageForm) GetSize() int {
	if pageForm.Size == 0 {
		return 10
	} else {
		return pageForm.Size
	}
}
