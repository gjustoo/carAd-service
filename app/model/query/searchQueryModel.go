package searchQueryModel

type SearchQuery struct {
	Id       int    `json:"id"`
	Search   string `json:"search"`
	Location string `json:"location"`
	MinPrice string `json:"min_price"`
	MaxPrice string `json:"max_price"`
	MinYear  string `json:"min_year"`
	MaxYear  string `json:"max_year"`
	MinKm    string `json:"min_km"`
	MaxKm    string `json:"max_km"`
}
