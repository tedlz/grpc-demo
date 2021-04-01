package transfer

import (
	"fmt"
	"strconv"
	"time"
)

// IdCard15To18 身份证号 15 位转 18 位
func IdCard15To18(oldIdCard string) (newIdCard string) {
	if len(oldIdCard) != 15 {
		newIdCard = oldIdCard
		return
	}
	iS := 0

	// 加权因子常数
	iW := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验码常数
	lastCode := "10X98765432"

	// 先在 6、7 位加上 19（不考虑 18XX 年的寿星，20XX 的应该都已经是 18 位了）
	newIdCard = oldIdCard[:6]
	newIdCard += "19"
	newIdCard += oldIdCard[6 : 6+9]
	// 进行加权求和
	for i := 0; i < 17; i++ {
		s, _ := strconv.Atoi(newIdCard[i : i+1])
		iS += s * iW[i]
	}

	// 取模运算，得到模值
	iY := iS % 11

	// 从 LastCode 中取得以模为索引号的值，加到身份证的最后一位，即为新身份证号
	newIdCard += lastCode[iY : iY+1]

	return newIdCard
}

// GetBornInfo 获取出生日期信息
func GetBornInfo(idCard string) (bornDate string, bornTime int64) {
	bornDate, bornTime = "0000-00-00 00:00:00", 0
	if len(idCard) != 18 {
		return
	}
	t, err := time.ParseInLocation("20060102", idCard[6:14], time.Local)
	if err != nil {
		return
	}
	bornDate, bornTime = t.Format("2006-01-02 15:04:05"), t.Unix()
	return
}

// GetAgeByIdCard 根据身份证号计算年龄
func GetAgeByIdCard(idCard string) (age int) {
	// 判断
	if length := len(idCard); length == 15 {
		// 如果是 15 位身份证，转 18 位计算
		idCard = IdCard15To18(idCard)
	} else if length == 18 {
		// 不处理
	} else {
		return
	}

	// 获取身份证年月日
	date := idCard[6:14]
	fmt.Println(date)
	t, err := time.ParseInLocation("20060102", date, time.Local)
	if err != nil {
		return
	}
	y, m, d := t.Date()

	// 获取当前年月日
	nowY, nowM, nowD := time.Now().Date()

	// 计算
	age = nowY - y
	if mo, nowMo := int(m), int(nowM); nowMo < mo ||
		(nowMo == mo &&
			nowD < d) {
		age--
	}
	return
}
