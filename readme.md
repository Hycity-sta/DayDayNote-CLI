# daydaynote

daydaynote 是一个用于记录日常生活的 Go 命令行工具，使用 [Cobra](https://github.com/spf13/cobra) 搭建。

## 可用命令

```bash
daydaynote version
daydaynote --help
daydaynote lang zh
daydaynote add 吃了个汉堡
daydaynote list
daydaynote list --date=2026/4/14
daydaynote delete 1
daydaynote delete 1 --date=2026-4-14
daydaynote edit 1 吃了两个汉堡
daydaynote edit 1 吃了两个汉堡 --date=2026/4/14
```

`list` 默认只查看当天内容；加上 `--date` 后，可以查看指定日期的数据。
日期支持 `YYYY/M/D` 和 `YYYY-MM-DD` 两种格式。

如果你还在开发调试阶段，也可以把上面的 `daydaynote` 替换成 `go run .` 使用。

## 构建

```bash
go build -o bin/daydaynote.exe
```

## 存储

默认使用 `exe` 同级 `data/年/月.jsonl`，每一行都是一个 JSON 对象，适合后续按行追加和逐行读取。

## 语言文件

语言资源通过 `embed` 打包在 `i18n/locales/en.json` 和 `i18n/locales/zh.json` 中，启动时会自动加载。

## 开发帮助

见`attention.md`文件。

## todo

- 将command里面的Run/RunE抽离出来，而不是嵌套一大坨函数
