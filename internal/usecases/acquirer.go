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
	Parse(content string) (*entities.AcquirerDetail, error)
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

func (uc *Acquirer) Parse(content string) (*entities.AcquirerDetail, error) {
	var detail entities.AcquirerDetail
	for len(content) > 0 {
		data, err := uc.dataUsecase.Parse(content)
		if err != nil {
			return nil, err
		}

		switch data.Tag {
		case uc.siteTag:
			detail.Site = *data
		case uc.mpanTag:
			detail.MPAN = *data
		case uc.terminalIDTag:
			detail.TerminalID = *data
		case uc.categoryTag:
			detail.Category = *data
		}

		content = content[4+len(data.Content):]
	}

	return &detail, nil
}
