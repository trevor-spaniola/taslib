package converts

import (
	"fmt"
	"strings"
)

func IntArrayToString(a []int, delim string) string {
	//return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
