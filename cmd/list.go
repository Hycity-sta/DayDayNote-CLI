package cmd

import (
	"fmt"
	"strings"
	"time"

	"daydaynote/i18n"
	"daydaynote/storage"
	"daydaynote/utils"

	"github.com/spf13/cobra"
)

// list 命令
func List() *cobra.Command {
	var dateValue string

	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdListUse),
		Short: i18n.T(i18n.CmdListShort),
		Args:  utils.NoArgs(),
		RunE: func(cmd *cobra.Command, args []string) error {
			targetDate, err := utils.ResolveDate(dateValue)
			if err != nil {
				return err
			}

			store := storage.StoreForDate(targetDate)
			records, err := store.List()
			if err != nil {
				return err
			}

			dayRecords := filterRecordsByDate(records, targetDate)
			if len(dayRecords) == 0 {
				fmt.Printf(i18n.T(i18n.MsgListEmpty), targetDate.Format("2006-01-02"))
				return nil
			}

			fmt.Println(i18n.T(i18n.MsgListHeader))
			for index, record := range dayRecords {
				fmt.Printf(
					i18n.T(i18n.MsgListRow),
					index+1,
					record.CreatedAt.Local().Format("15:04:05"),
					compactContent(record.Content),
				)
			}

			return nil
		},
	}

	command.Flags().StringVar(&dateValue, "date", "", i18n.T(i18n.FlagDate))
	return command
}

// 把超长内容压成单行，避免列表输出被换行打散
func compactContent(content string) string {
	content = strings.ReplaceAll(content, "\r\n", " ")
	content = strings.ReplaceAll(content, "\n", " ")
	return strings.TrimSpace(content)
}

// 从整个月的数据里筛出指定日期的记录
func filterRecordsByDate(records []storage.Record, targetDate time.Time) []storage.Record {
	filtered := make([]storage.Record, 0)
	for _, record := range records {
		recordDate := record.CreatedAt.Local()
		if utils.SameDate(recordDate, targetDate) {
			filtered = append(filtered, record)
		}
	}
	return filtered
}
