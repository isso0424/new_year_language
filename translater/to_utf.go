package translater

import (
	"strconv"
	"strings"
)

func translateToJapanese(raw string) string {
	result := strings.Replace(raw, "🎍", "0", -1)
	result = strings.Replace(result, "🍱", "1", -1)
	tmp := ""
	parsed := ""
	for index, char := range result {
		tmp += string(char)
		if ((index + 1) % 36 == 0) {
			tmpInt, _ := strconv.ParseInt(tmp, 2, 64)
			parsed += string(tmpInt)
			tmp = ""
		}
	}

	return parsed
}
