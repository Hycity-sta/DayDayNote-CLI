package main

import (
	"fmt"
	"os"
	"strings"

	"daydaynote/cmd"
	"daydaynote/i18n"
	"github.com/spf13/cobra"
)

func main() {
	normalizeLangArg()
	i18n.SetLang(detectLang(os.Args[1:]))

	rootCmd := &cobra.Command{
		Use:   "daydaynote",
		Short: i18n.T("root.short"),
		Long:  i18n.T("root.long"),
	}

	var lang string
	rootCmd.PersistentFlags().StringVar(&lang, "lang", "en", i18n.T("flag.lang"))

	rootCmd.AddCommand(cmd.NewVersionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func detectLang(args []string) string {
	for i, arg := range args {
		if arg == "--lang" || arg == "-lang" {
			if i+1 < len(args) {
				return args[i+1]
			}
		}

		if value, ok := strings.CutPrefix(arg, "--lang="); ok {
			return value
		}

		if value, ok := strings.CutPrefix(arg, "-lang="); ok {
			return value
		}
	}

	return "en"
}

func normalizeLangArg() {
	for i, arg := range os.Args {
		if arg == "-lang" {
			os.Args[i] = "--lang"
			continue
		}

		if value, ok := strings.CutPrefix(arg, "-lang="); ok {
			os.Args[i] = "--lang=" + value
		}
	}
}
