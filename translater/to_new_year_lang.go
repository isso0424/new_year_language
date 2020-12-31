package translater

import (
	"fmt"
	"strings"
)

func translateToNewYearLang(raw string) string {
	result := ""
	for _, char := range raw {
		tmp := fmt.Sprintf("%b", char)
		length := len(tmp)
		for i := 36 - length; i != 0; i-- {
			result += "0"
		}
		result += tmp
	}
	result = strings.Replace(result, "0", "🎍", -1)
	result = strings.Replace(result, "1", "🍱", -1)

	return result
}
