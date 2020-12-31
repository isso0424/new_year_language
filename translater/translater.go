package translater

func Translate(raw string, isDecode bool) string {
	if (isDecode) {
		return translateToJapanese(raw)
	}

	return translateToNewYearLang(raw)
}
