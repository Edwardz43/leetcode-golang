package solution

import "bytes"

/*
Runtime: 14 ms, faster than 45.65% of Go online submissions for Integer to Roman.
Memory Usage: 3.2 MB, less than 63.05% of Go online submissions for Integer to Roman.
*/

var (
	t = map[int]string{
		0: "",
		1: "M",
		2: "MM",
		3: "MMM",
	}
	h = map[int]string{
		0: "",
		1: "C",
		2: "CC",
		3: "CCC",
		4: "CD",
		5: "D",
		6: "DC",
		7: "DCC",
		8: "DCCC",
		9: "CM",
	}
	d = map[int]string{
		0: "",
		1: "X",
		2: "XX",
		3: "XXX",
		4: "XL",
		5: "L",
		6: "LX",
		7: "LXX",
		8: "LXXX",
		9: "XC",
	}
	g = map[int]string{
		0: "",
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
	}
	decimals = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans   = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func IntToRoman(num int) string {
	return t[num/1000] + h[num%1000/100] + d[num%100/10] + g[num%10]
}
func IntToRoman2(num int) string {
	var res bytes.Buffer
	for i, s := range romans {
		for decimals[i] <= num {
			res.WriteString(s)
			num -= decimals[i]
		}
	}
	return res.String()
}
