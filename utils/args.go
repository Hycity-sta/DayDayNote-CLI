package utils

// 统一放命令参数个数校验，顺便把 Cobra 默认英文报错改成本项目的多语言提示
import (
	"fmt"

	"daydaynote/i18n"

	"github.com/spf13/cobra"
)

// 不接受任何位置参数
func NoArgs() cobra.PositionalArgs {
	return func(command *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}

		return fmt.Errorf(i18n.T(i18n.ErrArgsNoArgs), len(args))
	}
}

// 要求精确的参数个数
func ExactArgs(expected int) cobra.PositionalArgs {
	return func(command *cobra.Command, args []string) error {
		if len(args) == expected {
			return nil
		}

		return fmt.Errorf(i18n.T(i18n.ErrArgsExact), expected, len(args))
	}
}

// 要求至少传入指定个数的参数
func MinimumNArgs(minimum int) cobra.PositionalArgs {
	return func(command *cobra.Command, args []string) error {
		if len(args) >= minimum {
			return nil
		}

		return fmt.Errorf(i18n.T(i18n.ErrArgsMinimum), minimum, len(args))
	}
}
