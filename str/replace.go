package str

import (
	"strings"
)

func ENSymbol(raw string) string {
	raw = strings.Replace(raw, "，", ",", -1)
	raw = strings.Replace(raw, "（", "(", -1)
	raw = strings.Replace(raw, "）", ")", -1)
	raw = strings.Replace(raw, "：", ":", -1)
	raw = strings.Replace(raw, "。", ".", -1)
	return raw
}
