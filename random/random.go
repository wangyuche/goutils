package random

import (
	c_rand "crypto/rand"
	"math/big"
)

type RandStringContent string

const (
	EngCapContent   RandStringContent = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	EngLowContent   RandStringContent = "abcdefghijklmnopqrstuvwxyz"
	SpecialsContent RandStringContent = "~=+%^*/()[]{}/!@#$?|"
	DigitsContent   RandStringContent = "0123456789"
)

type RandStringMode int

const (
	EngCapMode   RandStringMode = 1
	EngLowMode   RandStringMode = 2
	SpecialsMode RandStringMode = 4
	DigitsMode   RandStringMode = 8
	AllMode      RandStringMode = 16
)

func RandString(length int, types RandStringMode) string {
	var content RandStringContent
	var mask int = 1
	t := int(types)
	for t > 0 {
		r := mask & t
		t = t ^ mask
		mask = mask << 1
		switch r {
		case int(EngCapMode):
			content += EngCapContent
		case int(EngLowMode):
			content += EngLowContent
		case int(SpecialsMode):
			content += SpecialsContent
		case int(DigitsMode):
			content += DigitsContent
		case int(AllMode):
			content += (EngCapContent + EngLowContent + SpecialsContent + DigitsContent)
		}
	}
	b := make([]byte, length)
	for i := range b {
		result, _ := c_rand.Int(c_rand.Reader, big.NewInt(int64(len([]byte(content)))))
		var _r int = int(result.Uint64())
		b[i] = content[_r]
	}
	return string(b)
}

func RandInt(length int, max int) []int {
	b := make([]int, 0)
	for i := 0; i < length; i++ {
		result, _ := c_rand.Int(c_rand.Reader, big.NewInt(int64(max)))
		b = append(b, int(result.Uint64()))
	}
	return b
}

func RandByWeight(length int, weight []int) []int {
	b := make([]int, 0)
	var sum int = 0
	for _, v := range weight {
		sum += v
	}
	for i := 0; i < length; i++ {
		result, _ := c_rand.Int(c_rand.Reader, big.NewInt(int64(sum)))
		b = append(b, int(result.Uint64()))
	}
	return b
}
