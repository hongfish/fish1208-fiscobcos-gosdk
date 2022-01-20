package gintool

import "math"

type Pager struct {
	Page      int `form:"page" json:"page"`
	PageSize  int `form:"pageSize" json:"pageSize"`
	Total     int `form:"total" json:"total"`
	PageCount int `form:"pageCount" json:"pageCount"`
	NumStart  int `form:"numStart" json:"numStart"`
}

func CreatePager(page, pageSize int) *Pager {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	pager := new(Pager)
	pager.Page = page
	pager.PageSize = pageSize
	pager.setNumStart()
	return pager
}

func (p *Pager) setNumStart() {
	p.NumStart = (p.Page - 1) * p.PageSize
}

func (p *Pager) SetTotal(total int) {
	p.Total = total
	p.PageCount = int(math.Ceil(float64(total)) / float64(p.PageSize))
}
