package entities

type Switching struct {
	Tag     string          `json:"tag"`
	Content string          `json:"content"`
	Data    string          `json:"data"`
	Detail  SwitchingDetail `json:"detail"`
}

type SwitchingDetail struct {
	Site     Data `json:"site"`
	NMID     Data `json:"nmid"`
	Category Data `json:"category"`
}
