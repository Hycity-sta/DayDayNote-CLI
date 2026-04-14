# daydaynote

这是一个使用 [Cobra](https://github.com/spf13/cobra) 搭建的 Go CLI 项目。

## 目录结构

- `main.go`：程序入口
- `cmd/`：每个 command 一个文件
- `storage/`：JSONL 存储层
- `bin/`：构建输出目录
- `data/`：运行时数据目录，位于 `exe` 同级，存放 JSONL
- `i18n/locales/`：内嵌语言文件目录

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
