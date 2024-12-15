package entities

type Acquirer struct {
	Tag     string         `json:"tag"`
	Content string         `json:"content"`
	Data    string         `json:"data"`
	Detail  AcquirerDetail `json:"detail"`
}

type AcquirerDetail struct {
	Site       ExtractData `json:"site"`
	MPAN       ExtractData `json:"mpan"`
	TerminalID ExtractData `json:"terminal_id"`
	Category   ExtractData `json:"category"`
}
