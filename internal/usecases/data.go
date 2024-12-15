package usecases

import (
	"fmt"
	"strconv"

	"github.com/fyvri/go-qris/internal/domain/entities"
)

type Data struct {
}

type DataInterface interface {
	Extract(codeString string) (*entities.ExtractData, error)
	ModifyContent(extractData entities.ExtractData, content string) entities.ExtractData
}

func NewData() DataInterface {
	return &Data{}
}

func (u *Data) Extract(codeString string) (*entities.ExtractData, error) {
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
	return &entities.ExtractData{
		Tag:     tag,
		Content: content,
		Data:    tag + lengthCode + content,
	}, nil
}

func (uc *Data) ModifyContent(extractData entities.ExtractData, content string) entities.ExtractData {
	length := len(content)
	data := extractData.Tag + fmt.Sprintf("%02d", length) + content

	return entities.ExtractData{
		Tag:     extractData.Tag,
		Content: content,
		Data:    data,
	}
}
