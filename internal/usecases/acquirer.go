package usecases

import (
	"github.com/fyvri/go-qris/internal/domain/entities"
)

type Acquirer struct {
	dataUsecase        DataInterface
	acquirerDetailTags *AcquirerDetailTags
}

type AcquirerDetailTags struct {
	Site       string
	MPAN       string
	TerminalID string
	Category   string
}

type AcquirerInterface interface {
	Parse(content string) (*entities.AcquirerDetail, error)
}

func NewAcquirer(dataUsecase DataInterface, acquirerDetailTags *AcquirerDetailTags) AcquirerInterface {
	return &Acquirer{
		dataUsecase:        dataUsecase,
		acquirerDetailTags: acquirerDetailTags,
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
		case uc.acquirerDetailTags.Site:
			detail.Site = *data
		case uc.acquirerDetailTags.MPAN:
			detail.MPAN = *data
		case uc.acquirerDetailTags.TerminalID:
			detail.TerminalID = *data
		case uc.acquirerDetailTags.Category:
			detail.Category = *data
		}

		content = content[4+len(data.Content):]
	}

	return &detail, nil
}
