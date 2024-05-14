package model

type Category struct {
	ID            int32       `json:"id"`
	Name          string      `json:"name"`
	Subcategories []*Category `json:"subcategories"`
}

type Language struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
