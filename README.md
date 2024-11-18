# 🌱 基本介绍
## 1.1 项目简介
- ⚠️ 身份证号码生成器是按身份证验证规则生成虚拟身份证号，非真实身份证，仅供测试使用，请勿用于非法用途。
- 本项目可以模拟生成二代身份证号，校验二代身份证号;根据身份证获取地址年龄性别等
## 1.2 身份证号码组成
参考[GB 11643—1999 公民身份号码](https://blog.csdn.net/chenlu5201314/article/details/90484093)
## 1.3 使用说明
- 本功能需要go版本>=1.22.0
## 1.4 关于地址
- 本项目使用`https://gitee.com/dcloud/opendb/tree/master/collection/opendb-city-china`中data.json地址，
- 1.0.3版本地址嵌入到代码，不再从data.json读取。若地址更新，可手动在area.go中更新areaJson变量的值
### 安装
```go
go get github.com/gjing1st/idcard
```

### 随机生成身份证号
```go
package main

import (
	"fmt"
	"github.com/gjing1st/id-card"
)

func main() {
	idCard := idcard.RandIdCard()
	fmt.Println(idCard)
}
```
### 解析身份证号
```go
package main

import (
	"fmt"
	"github.com/gjing1st/id-card"
)

func main() {
	res, err := idcard.ParseNumber("210124195408204663")
	fmt.Println(res, err)
}
```
### 根据地址，出生日期，性别生成身份证号
```go
package main

import (
	"fmt"
	"github.com/gjing1st/id-card"
)

func main() {
	idCard, err := idcard.GenerateIdCard("山东省潍坊市寿光市", "20220301", "男")
	fmt.Println(idCard, err)
}
```
### 校验身份证是否正确

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(idcard.CheckIdCard("230000198209171361"))
}

```