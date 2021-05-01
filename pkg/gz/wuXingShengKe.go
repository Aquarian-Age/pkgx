package gz

import (
	"fmt"
	"strings"
)

//返回五行生克的字符形式
func GetWXSKS(gz string) string {
	wx1, wx2, gan, zhi := gzWuXing(gz)
	wxs := []string{"木", "火", "土", "金", "水", "木", "火"}
	var w1 int
	for i := 0; i < 5; i++ {
		if strings.EqualFold(wx1, wxs[i]) {
			w1 = i
		}
	}
	var w2 int
	for j := 0; j < 5; j++ {
		if strings.EqualFold(wx2, wxs[j]) {
			w2 = j
			break
		}
	}
	var n int
	for x := 0; x < len(wxs); x++ {
		//前者生/克后者
		k := w1 + 2
		s := w1 + 1
		if strings.EqualFold(wx1, wx2) {
			n = 0
			break
		}
		if strings.EqualFold(wx2, wxs[k]) {
			n = -1
			break
		}
		if strings.EqualFold(wx2, wxs[s]) {
			n = 1
			break
		}
		//后者生/克前者
		rk := w2 + 2
		rs := w2 + 1
		if strings.EqualFold(wx1, wxs[rk]) {
			n = -2
			break
		}
		if strings.EqualFold(wx1, wxs[rs]) {
			n = 2
			break
		}
	}
	///
	var sks string
	switch n {
	case 0:
		sks = fmt.Sprintf("%s%s比和%s%s", gan, wx1, zhi, wx2)
	case 1:
		sks = fmt.Sprintf("%s%s生%s%s", gan, wx1, zhi, wx2)
	case -1:
		sks = fmt.Sprintf("%s%s克%s%s", gan, wx1, zhi, wx2)
	case 2:
		sks = fmt.Sprintf("%s%s生%s%s", zhi, wx2, gan, wx1)
	case -2:
		sks = fmt.Sprintf("%s%s克%s%s", zhi, wx2, gan, wx1)
	}
	return sks
}

//五行生克 木火土金水 临位相生隔位相克
//比和n=0 前者生后者n=1 前者克后者n=-1 后者生前者n=2 后者克前者n=-2
func Wxsk(gz string) int {
	wx1, wx2, _, _ := gzWuXing(gz)
	wxs := []string{"木", "火", "土", "金", "水", "木", "火"}
	var w1 int
	for i := 0; i < 5; i++ {
		if strings.EqualFold(wx1, wxs[i]) {
			w1 = i
		}
	}
	var w2 int
	for j := 0; j < 5; j++ {
		if strings.EqualFold(wx2, wxs[j]) {
			w2 = j
			break
		}
	}
	var n int
	for x := 0; x < len(wxs); x++ {
		//前者生/克后者
		k := w1 + 2
		s := w1 + 1
		if strings.EqualFold(wx1, wx2) {
			n = 0
			break
		}
		if strings.EqualFold(wx2, wxs[k]) {
			n = -1
			break
		}
		if strings.EqualFold(wx2, wxs[s]) {
			n = 1
			break
		}
		//后者生/克前者
		rk := w2 + 2
		rs := w2 + 1
		if strings.EqualFold(wx1, wxs[rk]) {
			n = -2
			break
		}
		if strings.EqualFold(wx1, wxs[rs]) {
			n = 2
			break
		}

	}
	return n
}

//传入干支 返回干五行属性 支五行属性
func gzWuXing(gz string) (string, string, string, string) {
	zhiWx := map[string]string{
		"亥": "水", "子": "水",
		"寅": "木", "卯": "木",
		"辰": "土", "未": "土", "戌": "土", "丑": "土",
		"巳": "火", "午": "火",
		"申": "金", "酉": "金",
	}
	ganWx := map[string]string{
		"甲": "木", "乙": "木",
		"丙": "火", "丁": "火",
		"戊": "土", "己": "土",
		"庚": "金", "辛": "金",
		"壬": "水", "癸": "水",
	}
	var zs, gs string   //干支五行
	var gan, zhi string //干支
	for kz, vz := range zhiWx {
		if strings.ContainsAny(gz, kz) {
			zs = vz
			zhi = kz
			break
		}
	}
	for kg, vg := range ganWx {
		if strings.ContainsAny(gz, kg) {
			gs = vg
			gan = kg
			break
		}
	}
	return gs, zs, gan, zhi
}