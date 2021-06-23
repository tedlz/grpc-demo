package check

import (
	"regexp"
	"strconv"
)

func IdCard(idCard string) bool {
	// 长度验证
	length := len(idCard)
	if length != 18 {
		return false
	}

	// 正则验证
	pattern := `(^\d{8}(0\d|10|11|12)([0-2]\d|30|31)\d{3}$)|(^\d{6}(18|19|20)\d{2}(0[1-9]|10|11|12)([0-2]\d|30|31)\d{3}(\d|X)$)`
	match, err := regexp.MatchString(pattern, idCard)
	if err != nil || !match {
		return false
	}

	// 校验码验证
	iS := 0
	// 加权因子常数
	iW := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验码常数
	lastCode := "10X98765432"

	// 进行加权求和
	for i := 0; i < 17; i++ {
		s, _ := strconv.Atoi(idCard[i : i+1])
		iS += s * iW[i]
	}

	// 取模运算，得到模值
	iY := iS % 11
	// fmt.Println(idCard[length-1:], lastCode[iY:iY+1])
	if idCard[length-1:] != lastCode[iY:iY+1] {
		return false
	}
	return true
}
