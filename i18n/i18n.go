package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"
)

var (
	// currentLang 保存当前正在使用的语言代码，默认使用英文。
	currentLang = "en"
	// messages 按语言缓存所有翻译内容，键为语言代码。
	messages = map[string]map[string]string{}
)

//go:embed locales/*.json
var localeFiles embed.FS

func init() {
	// 启动时一次性加载所有内置语言文件，避免运行时重复读文件。
	for _, lang := range []string{"en", "zh"} {
		path := filepath.ToSlash(filepath.Join("locales", lang+".json"))
		data, err := localeFiles.ReadFile(path)
		if err != nil {
			panic(fmt.Errorf("load locale %s: %w", lang, err))
		}

		dict := map[string]string{}
		if err := json.Unmarshal(data, &dict); err != nil {
			panic(fmt.Errorf("parse locale %s: %w", lang, err))
		}

		messages[lang] = dict
	}
}

// SetLang 切换当前语言；如果传入的语言不存在，则回退到英文。
func SetLang(lang string) {
	if _, ok := messages[lang]; ok {
		currentLang = lang
		return
	}

	currentLang = "en"
}

// Lang 返回当前生效的语言代码。
func Lang() string {
	return currentLang
}

// T 根据 key 读取当前语言文案，并在缺失时回退到英文，最后返回 key 本身。
func T(key string) string {
	if value, ok := messages[currentLang][key]; ok {
		return value
	}

	if value, ok := messages["en"][key]; ok {
		return value
	}

	return key
}
