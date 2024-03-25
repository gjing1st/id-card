// Path: kernel
// FileName: check_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/25$ 10:44$

package idcard

import (
	"fmt"
	"testing"
)

func TestGetCheckCode(t *testing.T) {
	code := getCheckCode("110101200007280708")
	fmt.Println(code)
	code = getCheckCode("110101200007281604")
	fmt.Println(code)
	code = getCheckCode("110101200007286405")
	fmt.Println(code)
	code = getCheckCode("11010120000728010X")
	fmt.Println(code)
	code = getCheckCode("110101200007281604")
	fmt.Println(code)
}

func TestRandAreaCode(t *testing.T) {
	fmt.Println(randAreaCode())
	area, err := getAreaName(randAreaCode())
	fmt.Println("area", area, "err=", err)
}

func TestParseData(t *testing.T) {
	res, err := ParseNumber("230000198209171361")
	fmt.Println("err=", err)
	fmt.Println(res)
}

func TestRandBirthday(t *testing.T) {
	birth := randBirthday()
	fmt.Println(birth)
}

func TestRandOrderCode(t *testing.T) {
	for range 80 {
		code := randOrderCode()
		fmt.Println(code)
	}
}

func TestRandNumber(t *testing.T) {
	id := RandIdCard()
	fmt.Println(id)
	res, _ := ParseNumber(id)
	fmt.Println(res)
}

func TestRandAreaName(t *testing.T) {
	fmt.Println(randAreaName())
}
func TestGenerateIdCard(t *testing.T) {
	for range 3 {
		id, err := GenerateIdCard(randAreaName(), randBirthday(), "女")
		fmt.Println("id=", id, "err=", err)
		res, err1 := ParseNumber(id)
		if err1 != nil {
			fmt.Println("id解析失败", err1)
			return
		}
		fmt.Println("res=", res)
	}
}

func TestCheckIdCard(t *testing.T) {
	fmt.Println(CheckIdCard("23000019820917136"))
}
