package gopassRuler

var defaultLanguage = "zh_cn"

const (
	LangZHCN = "zh_cn"
	LangENUS = "en_us"
	LangZHTR = "zh_tr"
)

func SetDefaultLang(lang string) {
	defaultLanguage = lang
}

func GetDefaultLang() string {
	return defaultLanguage
}
