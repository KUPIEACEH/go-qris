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
	Parse(content string) (*entities.SwitchingDetail, error)
}

func NewSwitching(dataUsecase DataInterface, siteTag string, nmidTag string, categoryTag string) SwitchingInterface {
	return &Switching{
		dataUsecase: dataUsecase,
		siteTag:     siteTag,
		nmidTag:     nmidTag,
		categoryTag: categoryTag,
	}
}

func (uc *Switching) Parse(content string) (*entities.SwitchingDetail, error) {
	var detail entities.SwitchingDetail
	for len(content) > 0 {
		data, err := uc.dataUsecase.Parse(content)
		if err != nil {
			return nil, err
		}

		switch data.Tag {
		case uc.siteTag:
			detail.Site = *data
		case uc.nmidTag:
			detail.NMID = *data
		case uc.categoryTag:
			detail.Category = *data
		}

		content = content[4+len(data.Content):]
	}

	return &detail, nil
}
