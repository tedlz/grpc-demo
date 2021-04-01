package transfer

import (
	"strings"
	"unicode/utf8"
)

// HideName 姓名隐私保护
// 规则：
// 姓名故意捣蛋只输一个字的，直接返回星号
// 姓名是两个字的，隐去名；
// 姓名大于两个字的，取首尾两个字，且除首尾外根据中间的字符数量补相同数量的星号
// * 此方法未考虑复姓、少数民族姓氏
func HideName(name string) string {
	if r, length := []rune(name), utf8.RuneCountInString(name); length == 1 {
		name = "*"
	} else if length == 2 {
		name = string(r[:1]) + "*"
	} else {
		number := length - 2
		name = string(r[:1]) + strings.Repeat("*", number) + string(r[length-1:])
	}
	return name
}
