package utils

type Paging struct {
	Page        int   `json:"page" form:"page"`
	Limit       int   `json:"limit" form:"limit"`
	Count       int   `json:"count" form:"count"`
	CurrentPage int   `json:"current_page" form:"current_page"`
	PerPage     int   `json:"per_page" form:"per_page"`
	Total       int64 `json:"total" form:"-"`
	TotalPage   int   `json:"total_page" form:"total_page"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}
