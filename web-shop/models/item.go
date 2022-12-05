package models

type Item struct {
	Id            int     `json:"id"`
	Item_group_id int     `json:"item_group_id"`
	Price         float64 `json:"price"`
	Image         string  `json:"image"`
	Feature       string  `json:"feature"`
}
