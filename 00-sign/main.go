package main

import (
	"fmt"
	"log"
	"time"
	"unsafe"

	"grpc-demo/00-sign/sign"
)

// Request 请求结构体
type Request struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Sex       int64  `json:"sex"`
	Null      string `json:"null"`
	Bool      bool   `json:"bool"`
	UTF8      string `json:"utf8"`
	Timestamp int64  `json:"_t"`
	Signature string `json:"_sign"`
}

func main() {
	fmt.Println()
	fmt.Println("time:", time.Now().Unix())

	// 系统，版本号
	os, ver := "1", "v1.5.2"
	// 根据系统和版本号生成 secretKey
	secretKey := sign.GetSecretKey(os, ver)

	fmt.Println("secretKey:", secretKey, "length:", len(secretKey))

	// 根据 secretKey 解析系统和版本号
	os, ver, err := sign.ParseSecretKey(secretKey)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("os:", os, "ver:", ver)

	// 模拟请求传参
	req := Request{
		ID:        1,     // 数字类型测试
		Name:      "Q",   // 字符串类型测试
		Sex:       0,     // 0 测试
		Null:      "",    // 空字符串测试
		Bool:      false, // 布尔值测试
		UTF8:      "汉语",  // UTF-8 编码测试
		Timestamp: 1614599723,
		Signature: sign.Base64Encode("a5651ff4a375f9c5fd13982cc89d9b39adfdb6b6"),
	}
	fmt.Println(unsafe.Sizeof(req))
	// 新建验证实例
	check := sign.NewCheck()
	// 把 struct 结构体转为 map[string]interface{}
	if err := check.GetBody(req); err != nil {
		log.Fatalln(err)
	}
	// 设置允许的超时时间
	check.SetTimeout(5 * time.Minute)
	// 验证是否超时
	if err := check.CheckTimeout(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("body:", check)

	// 根据不同的 os 和 ver 获取 secretKey
	key := sign.GetConfigSecretKey(os, ver)
	// 设置 secretKey
	check.SetSecretKey(key)
	// 用上一步设置的 secretKey 生成签名并与请求入参的签名验证
	// if err := check.CheckMd5Signature(); err != nil {
	// 	log.Fatalln("md5 err:", err)
	// }
	if err := check.CheckHmacSha1Signature(); err != nil {
		log.Fatalln("hmac_sha1 err:", err)
	}
}

// all output
// time: 1614570473
// secretKey: 8Zi9uPULhQtAIGi2unUlsRinfwUOs1i9 length: 32
// os: 1 ver: v1.5.2
// body: &{map[_t:1614570455 bool:false id:1 name:Q null: sex:0 utf8:汉语] a5651ff4a375f9c5fd13982cc89d9b39adfdb6b6 1614570455 300000000000 }
// query: _t=1614570455&bool=false&id=1&name=Q&null=&sex=0&utf8=%E6%B1%89%E8%AF%AD
// a5651ff4a375f9c5fd13982cc89d9b39adfdb6b6
