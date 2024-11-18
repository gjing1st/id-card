// Path: area
// FileName: info.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/25$ 9:29$

package idcard

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var (
	map1           map[string]string
	InverseAreaMap map[string]string
	AreaMap        map[string]string
)

func GetAddress(k string) (v string) {
	if len(k) != 6 {
		return
	}
	v3 := map1[k] //区
	var (
		v1 string //省
		v2 string //市
	)
	k1 := k[:2] + "0000"
	if k[2:] == "0000" {
		//一级编号为省，直接返回
		v = v3
		return
	} else if k[4:] == "00" {
		//二级编号为市级，获取所属省后返回
		v1 = map1[k1]

		v = v1 + v3 //省拼接原本的值返回
		return
	} else {
		//三级编号，获取省市后拼接第三级返回
		v1 = map1[k1]
		k2 := k[:4] + "00"
		v2 = map1[k2]
		v = v1 + v2 + v3
		return
	}
}

func init() {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 读取文件内容
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	type Area struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}
	var area []Area
	_ = json.Unmarshal(byteValue, &area)
	map1 = make(map[string]string, len(area))
	AreaMap = make(map[string]string, len(area))
	InverseAreaMap = make(map[string]string, len(area))
	for _, v := range area {
		map1[v.Code] = v.Name
	}
	for k, _ := range map1 {
		address := GetAddress(k)
		AreaMap[k] = address
		InverseAreaMap[address] = k
	}
}
