package entities

type Switching struct {
	Tag     string          `json:"tag"`
	Content string          `json:"content"`
	Data    string          `json:"data"`
	Detail  SwitchingDetail `json:"detail"`
}

type SwitchingDetail struct {
	Site     ExtractData `json:"site"`
	NMID     ExtractData `json:"nmid"`
	Category ExtractData `json:"category"`
}
