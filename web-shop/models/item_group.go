package models

type ItemGroup struct {
	Id          int    `json:"-"`
	Category_id int    `json:"category_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
