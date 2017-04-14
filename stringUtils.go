package stringUtils

import (
	"strings"
)

func Trim(str string) string {
	return strings.Trim(str, " ")
}
