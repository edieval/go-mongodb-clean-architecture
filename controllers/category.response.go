package controllers

type Publication struct {
	Status string
}

type CategoryInfo struct {
	Type        string
	Code        string
	Label       string
	Description string
	Publication Publication
}

type Contents struct {
}

type CategoryResponse struct {
	Category CategoryInfo
	Contents Contents
}
