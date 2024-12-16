package entities

type Acquirer struct {
	Tag     string         `json:"tag"`
	Content string         `json:"content"`
	Data    string         `json:"data"`
	Detail  AcquirerDetail `json:"detail"`
}

type AcquirerDetail struct {
	Site       Data `json:"site"`
	MPAN       Data `json:"mpan"`
	TerminalID Data `json:"terminal_id"`
	Category   Data `json:"category"`
}
