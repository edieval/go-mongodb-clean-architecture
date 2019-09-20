package models

type AttributeRule struct {
	AttributeCode   string   `bson:"attributeCode,omitempty"`
	Operator        string   `bson:"operator,omitempty"`
	AttributeValues []string `bson:"attributeValues,omitempty"`
}

type PopulationRules struct {
	Type           string          `bson:"type,omitempty"`
	ModelCodes     []string        `bson:"modelCodes,omitempty"`
	AttributeRules []AttributeRule `bson:"attributeRules,omitempty"`
}

type standardFilterFields struct {
	Sort                  string `bson:"sort,omitempty" json:"sort,omitempty"`
	NumberValuesToDisplay int    `bson:"numberValuesToDisplay,omitempty" json:"numberValuesToDisplay,omitempty"`
}

type toggleFilterFields struct {
}

type Filter struct {
	Code                 string               `bson:"code,omitempty"`
	Label                string               `bson:"label,omitempty"`
	Type                 string               `bson:"type,omitempty"`
	Description          string               `bson:"description,omitempty"`
	Path                 string               `bson:"path,omitempty"`
	StandardFilterFields standardFilterFields `bson:"standardFilterFields,omitempty" json:"standardFilterFields,omitempty"`
	ToggleFilterFields   toggleFilterFields   `bson:"toggleFilterFields,omitempty" json:"toggleFilterFields,omitempty"`
	IsExtended           bool
}

type Publication struct {
	Status string `bson:"status,omitempty"`
}

type CategoryModel struct {
	Code            string            `bson:"code,omitempty"`
	CategoryType    string            `bson:"categoryType,omitempty"`
	Label           string            `bson:"label,omitempty"`
	Description     string            `bson:"description,omitempty"`
	Publication     Publication       `bson:"publication,omitempty"`
	PopulationRules []PopulationRules `bson:"populationRules,omitempty"`
	Filters         []Filter          `bson:"filters,omitempty"`
}
