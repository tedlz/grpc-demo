package sign

import (
	"encoding/hex"
)

// hexEncode hex 编码
func hexEncode(str string) string {
	return hex.EncodeToString([]byte(str))
}

// hexDecode hex 解码
func hexDecode(src string) string {
	dst, _ := hex.DecodeString(src)
	return string(dst)
}
