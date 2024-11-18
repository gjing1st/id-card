// Path: area
// FileName: info_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/25$ 10:53$

package idcard

import (
	"fmt"
	"testing"
)

func TestGetName(t *testing.T) {
	fmt.Println(AreaMap["350203"])
}

func TestGetCode(t *testing.T) {
	fmt.Println(InverseAreaMap["江苏省南京市大厂区"])
}

func TestRandArea(t *testing.T) {
	i := 0
	for k := range AreaMap {
		fmt.Println(k)
		i++
		if i == 5 {
			return
		}
	}

}

func TestCode(t *testing.T) {
	fmt.Println(InverseAreaMap["北京市北京市西城区"])
}
