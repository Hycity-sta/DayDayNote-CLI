# daydaynote

这是一个使用 [Cobra](https://github.com/spf13/cobra) 搭建的 Go CLI 项目。

## 可用命令

```bash
go run . version
go run . hello
```

## 构建

```bash
go build -o bin/daydaynote.exe
```

## 存储

默认使用 `exe` 同级 `data/daydaynote.jsonl`，每一行都是一个 JSON 对象，适合后续按行追加和逐行读取。

## 语言文件

语言资源通过 `embed` 打包在 `i18n/locales/en.json` 和 `i18n/locales/zh.json` 中，启动时会自动加载。
