package sign

var config = map[string]map[string]string{}

// 系统类型
const (
	ios     = "1"
	android = "2"
	web     = "4"
)

// init 初始化系统和版本的 secretKey
func init() {
	config[ios] = map[string]string{
		"v1.5.2": "8Zi9uPULhQtAIGi2unUlsRinfwUOs1i9",
	}
	config[android] = map[string]string{
		"v1.5.2": "5XiXUwU6hNtqIeieupU3s7iqfVUGs5iP",
	}
}

// GetConfigSecretKey 根据系统和版本获取 secretKey
func GetConfigSecretKey(os, ver string) string {
	return config[os][ver]
}
