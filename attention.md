# attention

## 注释规范

这个项目中的 Go 注释统一使用中文，并以帮助理解代码为目的，不写重复代码字面的废话。

- 不使用 GoDoc 风格注释。
- 不写 `Version 返回 version 命令`、`LoadLanguage 读取语言配置` 这种以函数名开头的句式。
- 优先说明“为什么这样做”“这段流程在做什么”，而不是逐字翻译代码。
- 对函数说明用途、流程或关键行为。
- 对结构体和字段说明它表示什么数据。
- 对复杂代码块说明这一步为什么存在。
- 对简单且一眼能看懂的代码，可以不写注释。
- 注释保持简短直接，一条注释只表达一个重点。
- 多步骤逻辑可以拆成多条流程注释。
- 路径、存储规则、回退逻辑这类不明显的信息应明确写出来。

推荐写法：

```go
// lang 命令
func Lang() *cobra.Command { ... }

// 先更新当前进程里的语言，保证后续提示信息立即生效
i18n.SetLang(lang)
```

不推荐写法：

```go
// Lang 返回 lang 命令。
func Lang() *cobra.Command { ... }

// SetLang sets current language.
i18n.SetLang(lang)
```

## 项目结构

这个项目目前按下面的目录职责来组织：

- 目录结构图：

```text
daydaynote/
├─ main.go
├─ cmd/
│  ├─ lang.go
│  └─ version.go
├─ config/
│  └─ lang.go
├─ i18n/
│  ├─ i18n.go
│  └─ locales/
│     ├─ en.json
│     └─ zh.json
├─ storage/
│  ├─ note.go
│  └─ store.go
├─ bin/
│  ├─ daydaynote.exe
│  └─ lang.json
├─ data/
│  └─ 2026/
│     └─ 04.jsonl
├─ readme.md
└─ attention.md
```

- `main.go`：程序入口，负责初始化语言、组装根命令并启动 CLI。
- `cmd/`：命令层，每个命令单独放一个文件。
- `config/`：配置读取与保存逻辑，例如语言配置。
- `i18n/`：国际化能力和语言资源加载逻辑。
- `i18n/locales/`：中英文文案文件。
- `storage/`：数据存储层，负责 JSONL 文件的读写。
- `bin/`：构建产物目录，也会放 exe 同级运行时文件。
- `data/`：运行时数据目录，位于 exe 同级，按“年/月.jsonl”组织记录文件。

目录设计上遵循这些原则：

- 入口放在 `main.go`，不要把具体业务逻辑堆在入口文件里。
- 命令相关逻辑放在 `cmd/`，配置相关逻辑放在 `config/`，存储相关逻辑放在 `storage/`。
- 文案和国际化资源集中放在 `i18n/`，不要把提示文本分散写在各处。
- 与运行环境相关的落盘文件统一放在 exe 同级目录下，便于打包和分发。
