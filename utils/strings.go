package utils

import "github.com/samber/lo"

func RandomUppercase(length int) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	arr := lo.Samples(runes, length)
	return string(arr)
}
