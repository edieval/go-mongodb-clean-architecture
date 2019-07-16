package models

type PopulationRules struct {
	ModelCodes []string `bson:"modelCodes,omitempty"`
}

type CategoryModel struct {
	Code            string            `bson:"code,omitempty"`
	CategoryType    string            `bson:"categoryType,omitempty"`
	PopulationRules []PopulationRules `bson:"populationRules,omitempty"`
}
