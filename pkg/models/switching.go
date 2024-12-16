package models

type Switching struct {
	Tag     string
	Content string
	Data    string
	Detail  SwitchingDetail
}

type SwitchingDetail struct {
	Site     Data
	NMID     Data
	Category Data
}
