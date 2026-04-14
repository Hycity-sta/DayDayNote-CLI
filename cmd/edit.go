package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"daydaynote/i18n"
	"daydaynote/storage"
	"daydaynote/utils"

	"github.com/spf13/cobra"
)

// edit 命令
func Edit() *cobra.Command {
	var dateValue string

	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdEditUse),
		Short: i18n.T(i18n.CmdEditShort),
		Args:  utils.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			targetDate, err := utils.ResolveDate(dateValue)
			if err != nil {
				return err
			}

			index, err := utils.ParsePositiveIndex(args[0])
			if err != nil {
				return err
			}

			content := strings.TrimSpace(strings.Join(args[1:], " "))
			if content == "" {
				return errors.New(i18n.T(i18n.ErrEditEmpty))
			}

			store := storage.StoreForDate(targetDate)
			records, err := store.List()
			if err != nil {
				return err
			}

			matchedPosition := 0
			editPosition := -1
			for position, record := range records {
				if !utils.SameDate(record.CreatedAt.Local(), targetDate) {
					continue
				}

				matchedPosition++
				if matchedPosition == index {
					editPosition = position
					break
				}
			}

			if editPosition < 0 {
				return fmt.Errorf(i18n.T(i18n.ErrEditIndexNotFound), targetDate.Format("2006-01-02"), index)
			}

			records[editPosition].Title = content
			records[editPosition].Content = content
			records[editPosition].UpdatedAt = time.Now()

			if err := store.Replace(records); err != nil {
				return err
			}

			fmt.Printf(i18n.T(i18n.MsgEditSuccess), index, targetDate.Format("2006-01-02"))
			return nil
		},
	}

	command.Flags().StringVar(&dateValue, "date", "", i18n.T(i18n.FlagDate))
	return command
}
