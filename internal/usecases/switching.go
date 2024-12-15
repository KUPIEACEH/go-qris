package usecases

import (
	"github.com/fyvri/go-qris/internal/domain/entities"
)

type Switching struct {
	dataUsecase DataInterface
	siteTag     string
	nmidTag     string
	categoryTag string
}

type SwitchingInterface interface {
	Extract(content string) (*entities.SwitchingDetail, error)
}

func NewSwitching(dataUsecase DataInterface, siteTag string, nmidTag string, categoryTag string) SwitchingInterface {
	return &Switching{
		dataUsecase: dataUsecase,
		siteTag:     siteTag,
		nmidTag:     nmidTag,
		categoryTag: categoryTag,
	}
}

func (uc *Switching) Extract(content string) (*entities.SwitchingDetail, error) {
	var detail entities.SwitchingDetail
	for len(content) > 0 {
		extractData, err := uc.dataUsecase.Extract(content)
		if err != nil {
			return nil, err
		}

		switch extractData.Tag {
		case uc.siteTag:
			detail.Site = *extractData
		case uc.nmidTag:
			detail.NMID = *extractData
		case uc.categoryTag:
			detail.Category = *extractData
		}

		content = content[4+len(extractData.Content):]
	}

	return &detail, nil
}
