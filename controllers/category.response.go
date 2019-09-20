package controllers

type Publication struct {
	Status string `json:"status,omitempty"`
}

type CategoryInfo struct {
	Type        string      `json:"type,omitempty"`
	Code        string      `json:"code,omitempty"`
	Label       string      `json:"label,omitempty"`
	Description string      `json:"description,omitempty"`
	Publication Publication `json:"publication,omitempty"`
}

type Content struct {
	Code string `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
}

type FilterValue struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Total int64  `json:"total,omitempty"`
}

type Filter struct {
	Type   string        `json:"type,omitempty"`
	Code   string        `json:"code,omitempty"`
	Label  string        `json:"label,omitempty"`
	Values []FilterValue `json:"values,omitempty"`
	Min    float64       `json:"min,omitempty"`
	Max    float64       `json:"max,omitempty"`
}

type Contents struct {
	Offset  int       `json:"offset,omitempty"`
	Limit   int       `json:"limit,omitempty"`
	Total   int64     `json:"total,omitempty"`
	Context string    `json:"context,omitempty"`
	Values  []Content `json:"values,omitempty"`
	Filters []Filter  `json:"filters,omitempty"`
}

type CategoryResponse struct {
	Category CategoryInfo `json:"category"`
	Contents Contents     `json:"contents"`
}
