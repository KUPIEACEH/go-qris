package usecases

import (
	"github.com/fyvri/go-qris/internal/domain/entities"
)

type Acquirer struct {
	dataUsecase   DataInterface
	siteTag       string
	mpanTag       string
	terminalIDTag string
	categoryTag   string
}

type AcquirerInterface interface {
	Extract(content string) (*entities.AcquirerDetail, error)
}

func NewAcquirer(dataUsecase DataInterface, siteTag string, mpanTag string, terminalIDTag string, categoryTag string) AcquirerInterface {
	return &Acquirer{
		dataUsecase:   dataUsecase,
		siteTag:       siteTag,
		mpanTag:       mpanTag,
		terminalIDTag: terminalIDTag,
		categoryTag:   categoryTag,
	}
}

func (uc *Acquirer) Extract(content string) (*entities.AcquirerDetail, error) {
	var detail entities.AcquirerDetail
	for len(content) > 0 {
		extractData, err := uc.dataUsecase.Extract(content)
		if err != nil {
			return nil, err
		}

		switch extractData.Tag {
		case uc.siteTag:
			detail.Site = *extractData
		case uc.mpanTag:
			detail.MPAN = *extractData
		case uc.terminalIDTag:
			detail.TerminalID = *extractData
		case uc.categoryTag:
			detail.Category = *extractData
		}

		content = content[4+len(extractData.Content):]
	}

	return &detail, nil
}
