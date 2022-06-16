package solution

import "bytes"

var a = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func helper(index int, digits string, res []string, tmp *bytes.Buffer) {
	if index == len(digits) {
		res = append(res, tmp.String())
		return
	}
	num := digits[index]
	nums := a[string(num)]
	for i := 0; i < len(nums); i++ {
		tmp.WriteString(nums[i])
		helper(index+1, digits, res, tmp)
	}
}

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	if len(digits) == 1 {
		return a[digits]
	}
	var (
		res []string
		tmp bytes.Buffer
	)
	helper(0, digits, res, &tmp)
	return res
}
