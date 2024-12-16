package models

type Acquirer struct {
	Tag     string
	Content string
	Data    string
	Detail  AcquirerDetail
}

type AcquirerDetail struct {
	Site       Data
	MPAN       Data
	TerminalID Data
	Category   Data
}
