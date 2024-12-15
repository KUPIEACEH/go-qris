package usecases

import (
	"fmt"
	"strings"
)

type CRC16CCITT struct {
}

type CRC16CCITTInterface interface {
	GenerateCode(code string) string
}

func NewCRC16CCITT() CRC16CCITTInterface {
	return &CRC16CCITT{}
}

func (uc *CRC16CCITT) GenerateCode(code string) string {
	charCodeAt := func(s string, i int) int {
		return int(s[i])
	}

	crc := 0xFFFF
	for c := 0; c < len(code); c++ {
		crc ^= charCodeAt(code, c) << 8
		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc = crc << 1
			}
		}
	}

	hex := crc & 0xFFFF
	hexStr := strings.ToUpper(fmt.Sprintf("%X", hex))
	if len(hexStr) == 3 {
		hexStr = "0" + hexStr
	}

	return hexStr
}
