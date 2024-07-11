package request

import "test-task-go/internal/model"

type PaginationService struct {
	Limit       int             `form:"limit" json:"limit"`
	Page        int             `form:"page" json:"page"`
	SortField   string          `form:"sort_field" json:"sort_field,omitempty"`
	SortOrder   string          `form:"sort_order" json:"sort_order,omitempty"`
	FilterField string          `form:"filter_field" json:"filter_field,omitempty"`
	FilterValue string          `form:"filter_value" json:"filter_value,omitempty"`
	Rows        []model.Service `json:"rows"`

	sortFields   map[string]struct{}
	filterFields map[string]map[string]struct{}
}

func (p *PaginationService) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginationService) GetLimit() int {
	if p.Limit < 1 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *PaginationService) GetPage() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Page
}

func (p *PaginationService) GetSort() string {
	return p.SortField + " " + p.SortOrder
}

func (p *PaginationService) GetFilterField() string {
	return p.FilterField
}

func (p *PaginationService) GetFilterValue() string {
	return p.FilterValue
}

func (p *PaginationService) HasSort() bool {
	_, ok := p.sortFields[p.SortField]
	return (p.SortOrder == "asc" || p.SortOrder == "desc") && ok
}

func (p *PaginationService) HasFilter() bool {
	values, ok := p.filterFields[p.FilterField]
	if !ok {
		return false
	}
	_, ok = values[p.FilterValue]

	return ok
}

func (p *PaginationService) AddSort(field string) {
	if p.sortFields == nil {
		p.sortFields = make(map[string]struct{})
	}

	p.sortFields[field] = struct{}{}
}

func (p *PaginationService) AddFilter(field string, values []string) {
	if p.filterFields == nil {
		p.filterFields = make(map[string]map[string]struct{})
	}
	m := make(map[string]struct{})
	for _, v := range values {
		m[v] = struct{}{}
	}
	p.filterFields[field] = m
}

func (p *PaginationService) adjust() {
	p.Limit = p.GetLimit()
	p.Page = p.GetPage()
	if !p.HasSort() {
		p.SortField = ""
		p.SortOrder = ""

	}
	if !p.HasFilter() {
		p.FilterField = ""
		p.FilterValue = ""
	}
}

func (p *PaginationService) SetRows(rows []model.Service) {
	p.adjust()
	p.Rows = rows
}
