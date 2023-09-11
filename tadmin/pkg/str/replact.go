package str

import (
	"strings"
)

// ToENSymbol 把中文字符转换成英文字符
func ToENSymbol(raw string) string {
	raw = strings.Replace(raw, "，", ",", -1)
	raw = strings.Replace(raw, "（", "(", -1)
	raw = strings.Replace(raw, "）", ")", -1)
	raw = strings.Replace(raw, "：", ":", -1)
	raw = strings.Replace(raw, "。", ".", -1)
	return raw
}
