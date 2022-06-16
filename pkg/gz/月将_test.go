/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT
 */

package gz

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/pub"
	"strings"
	"testing"
)

func TestGetYueJiang(t *testing.T) {
	want := []struct {
		yj, name, ts string
	}{
		{"丑", "大吉", "2020-12-21"},
		{"子", "神后", "2021-01-20"},
		{"亥", "登明", "2021-02-18"},
		{"戌", "河魁", "2021-03-20"},
		{"酉", "从魁", "2021-04-20"},
		{"申", "传送", "2021-05-21"},
		{"未", "小吉", "2021-06-21"},
		{"午", "胜光", "2021-07-22"},
		{"巳", "太乙", "2021-08-23"},
		{"辰", "天罡", "2021-09-23"},
		{"卯", "太冲", "2021-10-23"},
		{"寅", "功曹", "2021-11-22"},
		{"丑", "大吉", "2021-12-21"},
		{"子", "神后", "2022-01-20"},
		{"亥", "登明", "2022-02-19"},
		{"戌", "河魁", "2022-03-20"},
		{"酉", "从魁", "2022-04-20"},
		{"申", "传送", "2022-05-21"},
		{"未", "小吉", "2022-06-21"},
		{"午", "胜光", "2022-07-23"},
		{"巳", "太乙", "2022-08-23"},
		{"辰", "天罡", "2022-09-23"},
		{"卯", "太冲", "2022-10-23"},
		{"寅", "功曹", "2022-11-22"},
		{"丑", "大吉", "2022-12-22"},
	}
	_, zqt := getJieArr(2021)
	for i := 0; i < len(zqt); i++ {
		obj := NewGanZhi(zqt[i].Year(), int(zqt[i].Month()), zqt[i].Day(), zqt[i].Hour())
		//yj := GetYueJiang(zqt[i].Year(), int(zqt[i].Month()), zqt[i].Day(), pub.GetZhiS(obj.Mgz))
		yj := obj.YueJiangStruct()
		if !strings.EqualFold(yj.Zhi, want[i].yj) &&
			strings.EqualFold(yj.Name, want[i].name) &&
			strings.EqualFold(yj.ZhongQiT, want[i].ts) {
			t.Errorf("func GetYueJiang(%d %d %d %s)=%v want %v\n",
				zqt[i].Year(), int(zqt[i].Month()), zqt[i].Day(), pub.GetZhiS(obj.Mgz), yj, want[i])
		}
	}
}
func TestYJ_TaiChongTianMa(t *testing.T) {
	yjo := &YJ{Zhi: "亥"} //辰巳午未申酉戌亥子丑寅卯
	yjo = &YJ{Zhi: "戌"}  //巳午未申酉戌亥子丑寅卯辰
	yjo = &YJ{Zhi: "子"}  //卯辰巳午未申酉戌亥子丑寅
	hz := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	for i := 0; i < len(hz); i++ {
		s := yjo.TaiChongTianMa(hz[i])
		fmt.Print(s)
	}
	////[亥 子 丑 寅 卯 辰 巳 午 未 申 酉 戌]
	////[寅 卯 辰 巳 午 未 申 酉 戌 亥 子 丑]
	//yjo := &YJ{Zhi: "亥"}
	//hgz := "丙寅" //午
	////[午 未 申 酉 戌 亥 子 丑 寅 卯 辰 巳]
	////[午 未 申 酉 戌 亥 子 丑 寅 卯 辰 巳]
	//yjo = &YJ{Zhi: "午"}
	//hgz = "甲午" //卯
	////[子 丑 寅 卯 辰 巳 午 未 申 酉 戌 亥]
	////[卯 辰 巳 午 未 申 酉 戌 亥 子 丑 寅]
	//yjo = &YJ{Zhi: "子"}
	//hgz = "丁卯" //午
	////[丑 寅 卯 辰 巳 午 未 申 酉 戌 亥 子]
	////[戌 亥 子 丑 寅 卯 辰 巳 午 未 申 酉]
	//yjo = &YJ{Zhi: "丑"}
	//hgz = "戊戌" //子
	//tm := yjo.TaiChongTianMa(hgz)
	//fmt.Println(tm)
}

func TestYJ_TianSanMen(t *testing.T) {
	yjo := &YJ{Zhi: "亥"}
	hz := "午"
	mao, you, wei := yjo.TianSanMen(hz)
	fmt.Println(mao, you, wei)

	s := yjo.TianSanMenStruct(hz).TianSanMenString()
	fmt.Println(s)
}
