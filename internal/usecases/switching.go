package usecases

import (
	"github.com/fyvri/go-qris/internal/domain/entities"
)

type Switching struct {
	dataUsecase         DataInterface
	switchingDetailTags *SwitchingDetailTags
}

type SwitchingDetailTags struct {
	Site     string
	NMID     string
	Category string
}

type SwitchingInterface interface {
	Parse(content string) (*entities.SwitchingDetail, error)
}

func NewSwitching(dataUsecase DataInterface, switchingDetailTags *SwitchingDetailTags) SwitchingInterface {
	return &Switching{
		dataUsecase:         dataUsecase,
		switchingDetailTags: switchingDetailTags,
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
		case uc.switchingDetailTags.Site:
			detail.Site = *data
		case uc.switchingDetailTags.NMID:
			detail.NMID = *data
		case uc.switchingDetailTags.Category:
			detail.Category = *data
		}

		content = content[4+len(data.Content):]
	}

	return &detail, nil
}
