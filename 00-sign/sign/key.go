package sign

import (
	"errors"
	"fmt"
	"strings"

	"github.com/speps/go-hashids"
)

const (
	// Salt 盐
	Salt = "test"
	// MinLength 生成 secretKey 的最小长度
	MinLength = 30
	// 用来生成的字符串格式（os,ver）
	Format = "%s,%s"
)

// GetSecretKey 根据 os 和 ver 生成唯一用来请求的 secretKey
func GetSecretKey(os, ver string) string {
	h := getHashID()
	str := fmt.Sprintf(Format, os, ver)
	e, _ := h.EncodeHex(hexEncode(str))
	return e
}

// ParseSecretKey 解析 secretKey，得到 os 和 ver
func ParseSecretKey(secretKey string) (os, ver string, err error) {
	h := getHashID()
	d, _ := h.DecodeHex(secretKey)
	decode := hexDecode(d)
	if decode == "" {
		return "", "", errors.New("no authorization")
	}
	s := strings.Split(decode, ",")
	return s[0], s[1], nil
}

// getHashID 获取 hashids.HashID 实例
func getHashID() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = Salt
	hd.MinLength = MinLength
	h, _ := hashids.NewWithData(hd)
	return h
}
