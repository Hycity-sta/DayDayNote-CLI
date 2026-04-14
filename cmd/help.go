package cmd

import (
	"fmt"
	"strings"

	"daydaynote/i18n"

	"github.com/spf13/cobra"
)

// 为整棵命令树应用本地化的 help 和 usage 文案
func ConfigureHelp(root *cobra.Command) {
	root.InitDefaultCompletionCmd()
	localizeCompletionCommand(root)
	root.SetUsageTemplate(localizedUsageTemplate())
	root.SetHelpTemplate(localizedHelpTemplate())
	root.SetHelpCommand(localizedHelpCommand(root))
	applyHelpFlagUsage(root)
}

// 递归初始化 help flag，并改写默认说明文本
func applyHelpFlagUsage(command *cobra.Command) {
	command.InitDefaultHelpFlag()

	if flag := command.Flags().Lookup("help"); flag != nil {
		flag.Usage = fmt.Sprintf(i18n.T(i18n.FlagHelp), commandDisplayName(command))
	}

	for _, subCommand := range command.Commands() {
		applyHelpFlagUsage(subCommand)
	}
}

// 返回给用户看的命令名，空值时回退到“当前命令”
func commandDisplayName(command *cobra.Command) string {
	name := command.DisplayName()
	if strings.TrimSpace(name) == "" {
		return i18n.T(i18n.HelpThisCommand)
	}

	return name
}

// 本地化后的 help 子命令
func localizedHelpCommand(root *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   i18n.T(i18n.CmdHelpUse),
		Short: i18n.T(i18n.CmdHelpShort),
		Long:  fmt.Sprintf(i18n.T(i18n.CmdHelpLong), root.DisplayName()),
		Args:  cobra.ArbitraryArgs,
		RunE: func(command *cobra.Command, args []string) error {
			target := root
			if len(args) > 0 {
				found, remaining, err := root.Find(args)
				if err != nil || found == nil || len(remaining) > 0 {
					return fmt.Errorf(i18n.T(i18n.ErrHelpUnknownTopic), strings.Join(args, " "))
				}
				target = found
			}

			return target.Help()
		},
	}
}

// 帮助页模板本身不包含固定英文标题，交给本地化后的 UsageString 负责
func localizedHelpTemplate() string {
	return `{{with or .Long .Short}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`
}

// 参考 Cobra 默认模板，把固定标题替换成项目里的多语言文案
func localizedUsageTemplate() string {
	return fmt.Sprintf(`%s{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

%s
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

%s
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

%s{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

%s{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

%s
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

%s
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

%s{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

%s{{end}}
`,
		i18n.T(i18n.HelpUsage),
		i18n.T(i18n.HelpAliases),
		i18n.T(i18n.HelpExamples),
		i18n.T(i18n.HelpAvailableCommands),
		i18n.T(i18n.HelpAdditionalCommands),
		i18n.T(i18n.HelpFlags),
		i18n.T(i18n.HelpGlobalFlags),
		i18n.T(i18n.HelpAdditionalTopics),
		i18n.T(i18n.HelpMoreInformation),
	)
}

// 把 Cobra 自动生成的 completion 命令标题改成当前语言
func localizeCompletionCommand(root *cobra.Command) {
	completion, _, err := root.Find([]string{"completion"})
	if err != nil || completion == nil || completion.Name() != "completion" {
		return
	}

	completion.Short = i18n.T(i18n.CmdCompletionShort)
}
