### 生成私钥
生成RSA私钥：`openssl genrsa -out server.key 2048`
> 生成RSA私钥，命令的最后一个参数，将指定生成密钥的位数，如果没有指定，默认512

生成ECC私钥：`openssl ecparam -genkey -name secp384r1 -out server.key`
> 生成ECC私钥，命令为椭圆曲线密钥参数生成及操作，本文中ECC曲线选择的是secp384r1

### 生成公钥
`openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650`
> openssl req：生成自签名证书，-new指生成证书请求、-sha256指使用sha256加密、-key指定私钥文件、-x509指输出证书、-days 3650为有效期

### 注意
```
Go v1.15 已经弃用了 Common Name，所以上述方式生成的证书需要加上环境变量 GODEBUG=x509ignoreCN=0 才能使用
```
