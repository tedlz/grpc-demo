package sign

import (
	"errors"
	"fmt"
	"math"
	"net/url"
	"sort"
	"time"
)

const (
	// KeySignature 签名参数名
	KeySignature = "_sign"
	// KeyTimestamp 时间戳参数名
	KeyTimestamp = "_t"
)

type check struct {
	// body 请求入参
	body map[string]interface{}
	// signature 签名
	signature string
	// timestamp 客户端请求时的时间戳
	timestamp int64
	// timeout 超时，默认 5 分钟
	timeout time.Duration
	// secretKey
	secretKey string
}

func NewCheck() *check {
	return &check{
		body:      make(map[string]interface{}),
		signature: "default",
		timestamp: time.Now().Unix(),
		timeout:   time.Minute * 5,
	}
}

func (c *check) GetBody(value interface{}) error {
	// if reflect.TypeOf(value).Kind() != reflect.Struct {
	// 	return errors.New("value must be a struct")
	// }

	// 通过反射把 struct 转 map[string]interface{}
	decode, err := StructToMap(value, "json")
	if err != nil {
		return err
	}

	// 通过 json encode 和 decode 的方式把 struct 转 map[string]interface{}
	// bts, err := json.Marshal(value)
	// if err != nil {
	// 	return err
	// }
	// var decode map[string]interface{}
	// d := json.NewDecoder(bytes.NewReader(bts))
	// d.UseNumber() // 防止 json.Unmarshal 把大数字转为科学计数法
	// err = d.Decode(&decode)
	// if err != nil {
	// 	return err
	// }

	for k, v := range decode {
		// KeyTimestamp 作为请求体参与加密
		if k == KeyTimestamp {
			// c.timestamp, _ = v.(json.Number).Int64() // 如果用 d.Decode(&decode) 需要配合这句
			c.timestamp = v.(int64)
		}
		// KeySignature 不参与加密
		if k == KeySignature {
			c.signature = Base64Decode(v.(string))
			continue
		}
		c.body[k] = v // 要加密的参数
	}
	return nil
}

// SetTimeout 设置超时时间
func (c *check) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// CheckTimeout 验证是否超时
func (c *check) CheckTimeout() error {
	timestamp := time.Unix(c.timestamp, 0)
	if math.Abs(time.Now().Sub(timestamp).Seconds()) > c.timeout.Seconds() {
		return errors.New("request timeout")
	}
	return nil
}

// SetSecretKey 设置 secretKey
func (c *check) SetSecretKey(secretKey string) {
	c.secretKey = secretKey
}

// GetSecretKey 获取 secretKey
func (c *check) GetSecretKey() string {
	return c.secretKey
}

// GetSignature 获取 signature
func (c *check) GetSignature() string {
	return c.signature
}

// GetUri 获取根据字典序拼接好的字符串
func (c *check) GetUri() string {
	var keys []string
	for k, _ := range c.body {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	values := url.Values{}
	for _, v := range keys {
		values.Add(v, fmt.Sprintf("%v", c.body[v]))
	}
	query := values.Encode()
	fmt.Println("query: " + query)
	return query
}

// CheckMd5Signature 校验 md5 签名
func (c *check) CheckMd5Signature() error {
	signature := GetMd5Signature(c.GetSecretKey(), c.GetUri())
	fmt.Println(signature)
	if signature != c.GetSignature() {
		return errors.New("invalid signature")
	}
	return nil
}

// CheckHmacSha1Signature 校验 hmac_sha1 签名
func (c *check) CheckHmacSha1Signature() error {
	signature := GetHmacSha1Signature(c.GetSecretKey(), c.GetUri())
	fmt.Println(signature)
	if signature != c.GetSignature() {
		return errors.New("invalid signature")
	}
	return nil
}
