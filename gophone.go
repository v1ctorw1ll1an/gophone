package gophone

import "fmt"

func FormatPhoneBr(number string, prefix string) string {
	n := number

	n = fmt.Sprintf("%s%s", prefix, n)

	return n
}
