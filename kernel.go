// Path: kernel
// FileName: check.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/25$ 10:36$

package idcard

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParseNumber
// @Description 解析身份证号
// @params id string
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 11:24
func ParseNumber(idCard string) (res ParseIdCard, err error) {
	// 身份证号正则表达式
	reg := `^[1-9]\d{5}(19\d{2}|20[0-2]\d)(0[1-9]|1[0-2])([0-2][1-9]|[1-3][0-1])\d{3}([0-9]|X)$`
	match, err1 := regexp.MatchString(reg, idCard)
	if err1 != nil {
		fmt.Println("身份证号校验出错：", err.Error())
		err = err1
		return
	}
	if !match {
		err = errors.New("身份证号校验失败")
		return
	}
	res.Number = idCard
	//获取地区
	res.Area, err = GetAreaName(idCard[:6])
	if err != nil {
		return
	}
	//生日
	res.Birthday = idCard[6:14]

	birthDate, _ := time.Parse("20060102", res.Birthday)
	ageDuration := time.Since(birthDate)

	res.Age = int(ageDuration.Hours() / 24 / 365)
	//校验码是否正确
	checkCode := GetCheckCode(idCard)
	if checkCode != string(idCard[17]) {
		res.CheckCode = false
	} else {
		res.CheckCode = true
	}
	//性别
	sexCode, _ := strconv.Atoi(string(idCard[16]))
	if sexCode%2 == 1 {
		res.Sex = "男"
	} else {
		res.Sex = "女"
	}
	return
}

// RandIdCard
// @Description 随机生成身份证号码
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 14:13
func RandIdCard() (idCard string) {
	//随机地址
	areaCode := RandAreaCode()
	birthday := RandBirthday()
	orderCode := RandOrderCode()
	sexCode := rand.IntN(10)
	idCard = areaCode + birthday + orderCode + strconv.Itoa(sexCode)
	idCard += GetCheckCode(idCard)
	return
}

// GenerateIdCard
// @Description 生成身份证号
// @params area string 省市区地址
// @params birthday string 生日20060102格式
// @params sex string 男/女
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 14:28
func GenerateIdCard(area, birthday, sex string) (idCard string, err error) {
	areaCode, err1 := GetAreaCode(area)
	if err1 != nil {
		err = err1
		return
	}
	birthday = strings.Replace(birthday, "-", "", -1)
	orderCode := RandOrderCode()
	sexCode := 0
	sexArr := []int{0, 2, 4, 6, 8}
	sexRandIndex := rand.IntN(5)
	if sex == "男" {
		sexArr = []int{1, 3, 5, 7, 9}
	}
	sexCode = sexArr[sexRandIndex]
	idCard = areaCode + birthday + orderCode + strconv.Itoa(sexCode)
	idCard += GetCheckCode(idCard)
	return
}

// GetCheckCode
// @Description 根据身份证号前17位生成最后一位校验码
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 11:05
func GetCheckCode(id string) (code string) {
	if len(id) < 17 {
		return
	}
	checkNum := 0
	for i := range 17 {
		idNum, _ := strconv.Atoi(string(id[i]))
		checkNum += ((1 << (17 - i)) % 11) * idNum
	}
	checkCode := (12 - (checkNum % 11)) % 11
	if checkCode == 10 {
		code = "X"
	} else {
		code = strconv.Itoa(checkCode)
	}
	return
}

// RandAreaCode
// @Description 随机获取一个地区号码
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 11:20
func RandAreaCode() string {
	for k := range AreaMap {
		return k
	}
	return ""
}

// RandAreaName
// @Description 随机地址
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 14:30
func RandAreaName() string {
	for k := range InverseAreaMap {
		return k
	}
	return ""
}

// GetAreaName
// @Description 获取地址码对应的地址
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 14:12
func GetAreaName(code string) (string, error) {
	if v, ok := AreaMap[code]; ok {
		return v, nil
	}
	return "", errors.New("地址码错误")
}

// GetAreaCode
// @Description 根据地址获取地址码
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 14:18
func GetAreaCode(name string) (string, error) {
	if v, ok := InverseAreaMap[name]; ok {
		return v, nil
	}
	return "", errors.New("地址错误")
}

// RandBirthday
// @Description 随机生成1900年1月1日至今天的某一天时间
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 13:53
func RandBirthday() string {
	startDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
	now := time.Now()
	diff := now.Sub(startDate)
	randDuration := rand.N(int64(diff))
	birthday := startDate.Add(time.Duration(randDuration))
	return birthday.Format("20060102")
}

// RandOrderCode
// @Description 随机生成顺序码(不包含性别)
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/25 14:12
func RandOrderCode() string {
	code := rand.IntN(88) + 11
	return strconv.Itoa(code)
}
