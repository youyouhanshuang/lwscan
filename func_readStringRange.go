package main

import (
	"strconv"
	"strings"
)

//处理传来的端口字符串列表,并完成拼接写入切片
func readPortString(Pstring string) []int {
	var targetPortList []int
	if strings.Contains(Pstring, ",") { //判断是否含有","
		pStringlist := strings.Split(Pstring, ",") //以","切割,并存入数组
		targetPortList = toAllPortlist(pStringlist)
		return targetPortList
	}
	targetPortList = readStringRange(Pstring)
	return targetPortList
}

//处理只含有"-"和单个端口字符串
func readStringRange(Ptsport string) []int {
	var shou string
	var wei string
	var listA []int
	if strings.Contains(Ptsport, "-") { //判断是否含有"-"
		for i := 0; i < len(Ptsport); i++ {
			if string(Ptsport[i]) == "-" { //识别"-"的位置
				shou = Ptsport[0:i]
				wei = Ptsport[i+1:]
			}
		}
		intShou, _ := strconv.Atoi(shou) //将首尾数字转为int
		intWei, _ := strconv.Atoi(wei)
		for j := intShou; j <= intWei; j++ {
			listA = append(listA, j) //写入切片
		}
		return (listA)
	}
	intPtsport, _ := strconv.Atoi(Ptsport)
	listA = append(listA, intPtsport)
	return (listA)
}

// func stringToInt(strList []string) (intList []int) {
// 	for i := 0; i < len(strList); i++ { //转为int型
// 		intList[i], _ = strconv.Atoi(strList[i])
// 	}
// 	return intList
// }

//{123,23,421,12,123-12,124,11-23}
//{ 0 ,1 ,2  ,3 ,4     ,5  ,6}

//将含"-"和","的端口列表整合成一个切片
func toAllPortlist(strPortList []string) []int {
	var allPortList []int
	for i := 0; i < len(strPortList); i++ {
		if strings.Contains(strPortList[i], "-") { //判断是否含有"-"
			intPortlist := readStringRange(strPortList[i])    //转为该范围切片
			allPortList = append(allPortList, intPortlist...) //附加在allPortList上
		} else {
			intPort, _ := strconv.Atoi(strPortList[i]) //无"-"时,直接将string转为int
			allPortList = append(allPortList, intPort) //放入切片
		}
	}
	return allPortList
}

// func stringToInt(strList []string) (intList []int) {
// 	for i := 0; i < len(strList); i++ { //转为int型切片
// 		intList[i], _ = strconv.Atoi(strList[i])
// 	}
// 	return intList
// }

// func main() {
// 	a := "10-100"         //[10,11,12,13,14....,99,100]
// 	b := "20"             //[20]
// 	c := "12,13,14,15-18" //[12,13,14,15,16,17,18]
// 	readPortString(a)
// 	readPortString(b)
// 	readPortString(c)
// }

//取出字典中端口对应的服务字符串
func takeServiceString(portNum string, inStr []string) string {
	var strServer string
	for {
		if portNum == inStr[0] {
			strServer = inStr[1]
			return strServer
		}
		if len(inStr) == 2 {
			strServer = "未知"
			return strServer
		}
		inStr = append(inStr[:0], inStr[2:]...)
	}
}
