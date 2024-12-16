package usecases

import (
	"fmt"
	"strconv"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

type Data struct {
}

type DataInterface interface {
	Parse(codeString string) (*entities.Data, error)
	ModifyContent(data *entities.Data, content string) *entities.Data
}

func NewData() DataInterface {
	return &Data{}
}

func (u *Data) Parse(codeString string) (*entities.Data, error) {
	if len(codeString) < 5 {
		return nil, fmt.Errorf("invalid format code")
	}

	tag := codeString[:2]
	lengthCode := codeString[2:4]
	length, err := strconv.Atoi(lengthCode)
	if err != nil {
		return nil, fmt.Errorf("invalid length format for tag %s: %s", tag, err)
	}

	if len(codeString) < 4+length {
		return nil, fmt.Errorf("invalid length for tag %s", tag)
	}

	content := codeString[4 : 4+length]
	return &entities.Data{
		Tag:     tag,
		Content: content,
		Data:    tag + lengthCode + content,
	}, nil
}

func (uc *Data) ModifyContent(data *entities.Data, content string) *entities.Data {
	length := len(content)
	if length < 1 {
		return &entities.Data{}
	}

	return &entities.Data{
		Tag:     data.Tag,
		Content: content,
		Data:    data.Tag + fmt.Sprintf("%02d", length) + content,
	}
}
