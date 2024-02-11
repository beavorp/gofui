package node

func RawCode(attrs, content string) string {
	return "<code " + attrs + ">" + content + "</code>"
}
