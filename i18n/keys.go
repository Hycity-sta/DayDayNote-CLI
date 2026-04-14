package i18n

// 统一设置所有翻译 key，方便在编辑器里获得自动补全和重命名能力
const (
	AppName = "app.name"

	RootShort = "root.short"
	RootLong  = "root.long"

	CmdAddUse       = "cmd.add.use"
	CmdAddShort     = "cmd.add.short"
	CmdListUse      = "cmd.list.use"
	CmdListShort    = "cmd.list.short"
	CmdDeleteUse    = "cmd.delete.use"
	CmdDeleteShort  = "cmd.delete.short"
	CmdEditUse      = "cmd.edit.use"
	CmdEditShort    = "cmd.edit.short"
	CmdHelpUse      = "cmd.help.use"
	CmdHelpShort    = "cmd.help.short"
	CmdLangUse      = "cmd.lang.use"
	CmdLangShort    = "cmd.lang.short"
	CmdVersionUse   = "cmd.version.use"
	CmdVersionShort = "cmd.version.short"

	ErrAddEmpty            = "err.add.empty"
	ErrArgsNoArgs          = "err.args.no_args"
	ErrArgsExact           = "err.args.exact"
	ErrArgsMinimum         = "err.args.minimum"
	ErrEditEmpty           = "err.edit.empty"
	ErrEditIndexNotFound   = "err.edit.index_not_found"
	ErrDateInvalid         = "err.date.invalid"
	ErrIndexInvalid        = "err.index.invalid"
	ErrDeleteIndexNotFound = "err.delete.index_not_found"
	ErrLangUnsupported     = "err.lang.unsupported"

	FlagDate = "flag.date"

	MsgAddSaved      = "msg.add.saved"
	MsgEditSuccess   = "msg.edit.success"
	MsgListEmpty     = "msg.list.empty"
	MsgListHeader    = "msg.list.header"
	MsgListRow       = "msg.list.row"
	MsgDeleteSuccess = "msg.delete.success"
	MsgLangSet       = "msg.lang.set"
	MsgVersion       = "msg.version"
)
