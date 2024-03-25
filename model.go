// Path: idcard
// FileName: model.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/25$ 11:29$

package idcard

// ParseIdCard 解析身份证信息
type ParseIdCard struct {
	Number    string `json:"number"`
	Area      string `json:"area"`
	Birthday  string `json:"birthday"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	CheckCode bool   `json:"check_code"`
}
