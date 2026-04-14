# daydaynote

这是一个使用 [Cobra](https://github.com/spf13/cobra) 搭建的 Go CLI 项目。

## 目录结构

- `main.go`：程序入口
- `cmd/`：每个 command 一个文件
- `storage/`：JSONL 存储层
- `bin/data/`：默认数据文件目录

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

默认使用 `bin/data/daydaynote.jsonl`，每一行都是一个 JSON 对象，适合后续按行追加和逐行读取。
