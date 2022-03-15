/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT
 */

package gz

import (
	"reflect"
	"testing"
)

func TestXianChi(t *testing.T) {
	want := map[string]string{
		"子": "酉", "丑": "午", "寅": "卯", "卯": "子", "辰": "酉", "巳": "午",
		"午": "卯", "未": "子", "申": "酉", "酉": "午", "戌": "卯", "亥": "子",
	}
	var smap = make(map[string]string)
	for i := 1; i < len(Zhis); i++ {
		s := XianChi(Zhis[i])
		smap[Zhis[i]] = s
	}
	if !reflect.DeepEqual(smap, want) {
		t.Errorf("func XianChi()=%v want:%v", smap, want)
	}
}
