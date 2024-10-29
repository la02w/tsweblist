package utils

import (
	"math/rand"
)

const (
	LowerBytes  = "abcdefghijklmnopqrstuvwxyz"
	UpperBytes  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes = "0123456789"
	allBytes    = LowerBytes + UpperBytes + numberBytes
)

func GeneratePassword() string {
	// 创建密码
	password := make([]byte, 6)

	// 确保密码至少包含一个小写字母、一个大写字母和一个数字
	password[0] = LowerBytes[rand.Intn(len(LowerBytes))]
	password[1] = UpperBytes[rand.Intn(len(UpperBytes))]
	password[2] = numberBytes[rand.Intn(len(numberBytes))]

	// 如果密码长度大于 3，填充剩余的字符
	for i := 3; i < 6; i++ {
		password[i] = allBytes[rand.Intn(len(allBytes))]
	}

	// 打乱密码中的字符
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}
