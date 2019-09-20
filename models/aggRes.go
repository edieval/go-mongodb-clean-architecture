package models

type Number struct {
	Number string `json:"$numberInt"`
}

type Double struct {
	Double string `json:"$numberDouble"`
}

type Total struct {
	Total Number `json:"total,omitempty"`
}

type Product struct {
	Id          string `json:"_id"`
	SortedField Number `json:"sortedField"`
	Type        string `json:"type"`
}

type FacetResult struct {
	Id    string `json:"_id"`
	Count Number `json:"count"`
}

type FacetRangeResult struct {
	Id  string `json:"_id"`
	Min Double `json:"min"`
	Max Double `json:"max"`
}
