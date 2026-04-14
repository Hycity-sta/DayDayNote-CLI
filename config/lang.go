package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	// 英文语言代码
	LanguageEnglish = "en"
	// 中文语言代码
	LanguageChinese = "zh"
)

// exe 同级 lang.json 的文件结构
type langConfig struct {
	Lang string `json:"lang"`
}

// 读取 exe 同级目录中的语言配置；
// 文件不存在或内容无效时返回英文
func LoadLanguage() string {
	data, err := os.ReadFile(langFilePath())
	if err != nil {
		return LanguageEnglish
	}

	var cfg langConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return LanguageEnglish
	}

	if !IsSupportedLanguage(cfg.Lang) {
		return LanguageEnglish
	}

	return cfg.Lang
}

// 将语言配置写入 exe 同级目录
func SaveLanguage(lang string) error {
	if !IsSupportedLanguage(lang) {
		return os.ErrInvalid
	}

	data, err := json.MarshalIndent(langConfig{Lang: lang}, "", "  ")
	if err != nil {
		return err
	}

	data = append(data, '\n')
	return os.WriteFile(langFilePath(), data, 0o644)
}

// 判断是否为支持的语言
func IsSupportedLanguage(lang string) bool {
	return lang == LanguageEnglish || lang == LanguageChinese
}

// 返回语言配置文件路径，位置固定在 exe 同级目录；
// 当无法获取可执行文件路径时，回退到当前工作目录下的 lang.json
func langFilePath() string {
	exe, err := os.Executable()
	if err != nil {
		return "lang.json"
	}

	return filepath.Join(filepath.Dir(exe), "lang.json")
}
