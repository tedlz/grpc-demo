package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

// GetMd5Signature 生成 md5 的签名
func GetMd5Signature(secretKey, uri string) string {
	h := md5.New()
	io.WriteString(h, secretKey+uri)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GetHmacSha1Signature 生成 hmac_sha1 签名
func GetHmacSha1Signature(secretKey, uri string) string {
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(uri))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Base64Encode base64 编码
func Base64Encode(value string) string {
	return base64.URLEncoding.EncodeToString([]byte(value))
}

// Base64Decode base64 解码
func Base64Decode(value string) string {
	b, _ := base64.URLEncoding.DecodeString(value)
	return string(b)
}
