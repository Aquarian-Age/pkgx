/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT
 */

package pub

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// GetGanS 取天干 传干支 返回干
func GetGanS(gz string) string {
	gan := [11]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	var g string
	for i := 0; i < len(gan); i++ {
		if strings.ContainsAny(gz, gan[i]) {
			g = gan[i]
			break
		}
	}
	return g
}

// GetZhiS 取地支 传干支返回支
func GetZhiS(gz string) string {
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var z string
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(gz, zhi[i]) {
			z = zhi[i]
			break
		}
	}
	return z
}

// SortArr 顺序排地支 传入对应的地支 原地支数组 返回排序后的地支数组
func SortArr(zhi string, zhiArr []string) []string {
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(zhi, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	return zhiArr
}

// ReSortArr 逆序排地支
func ReSortArr(zhi string, zhiArr []string) []string {
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(zhi, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	//
	head := zhiArr[:1]
	end := zhiArr[1:]
	var rArr []string
	for i := len(end) - 1; i >= 0; i-- {
		rArr = append(rArr, end[i])
	}
	rArr = append(head, rArr...)
	return rArr
}

// ReName 顺序不变　把arr数组中source名字改为source + add 如果没有符合条件的返回原值
func ReName(arr []string, source, add string) []string {
	var newArr []string
	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(arr[i], source) {
			source = source + add
			head := arr[:i+1]
			head = head[:len(head)-1]
			head = append(head, source)
			end := arr[i+1:]
			newArr = append(head, end...)
		} else {
			newArr = arr
		}
	}
	return newArr
}

// ReNameS 正则替换 把s的dels元素替换成r
func ReNameS(s, dels string, r string) string {
	rs, err := regexp.Compile(dels)
	if err != nil {
		log.Fatal(err)
	}
	return rs.ReplaceAllString(s, r)
}

// SortArrInt 顺排 int数组 排序后使得index在首位 返回排序后的数组 如果index不在数组中 返回错误
func SortArrInt(index int, arr []int) ([]int, error) {
	var err error
	for i := 0; i < len(arr); i++ {
		if index == arr[i] {
			head := arr[i:]
			end := arr[:i]
			arr = append(head, end...)
			return arr, nil
		} else if index != arr[i] {
			err = fmt.Errorf("%d不在数组%d中", index, arr)
		}
	}
	return arr, err
}

// SortInt 正向排序
func SortInt(index int, arr []int) []int {
	var xarr []int
	for i := 0; i < len(arr); i++ {
		if index == arr[i] {
			head := arr[i:]
			end := arr[:i]
			xarr = append(head, end...)
			break
		}
	}
	return xarr
}

// ReSortArrInt 逆向排序
func ReSortArrInt(index int, arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if index == arr[i] {
			head := arr[i:]
			end := arr[:i]
			arr = append(head, end...)
			break
		}
	}
	head := arr[:1]
	end := arr[1:]
	var rArr []int
	for i := len(end) - 1; i >= 0; i-- {
		rArr = append(rArr, end[i])
	}
	rArr = append(head, rArr...)
	return rArr
}

// WeekName 周几 0,7周日　1:周一
func WeekName(n int) (w string) {
	switch n {
	case 0, 7:
		w = "周日"
	case 1:
		w = "周一"
	case 2:
		w = "周二"
	case 3:
		w = "周三"
	case 4:
		w = "周四"
	case 5:
		w = "周五"
	case 6:
		w = "周六"
	}
	return
}

// DelElement 去除切片空元素
func DelElement(s []string) []string {
	news := []string{}
	for _, v := range s {
		if len(v) > 0 {
			news = append(news, v)
		}
	}
	return news
}
