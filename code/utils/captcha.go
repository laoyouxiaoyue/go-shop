package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func RandomDigitCaptcha(len int) (string, error) {
	if len <= 0 {
		return "", fmt.Errorf("长度必须大于0")
	}

	digits := "0123456789"
	result := make([]byte, len)

	// 第一位：1~9
	first, err := rand.Int(rand.Reader, big.NewInt(9))
	if err != nil {
		return "", err
	}
	result[0] = digits[first.Int64()+1] // +1 避免 0

	// 其余位：0~9
	for i := 1; i < len; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		result[i] = digits[num.Int64()]
	}

	return string(result), nil
}
