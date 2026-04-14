package cmd

import (
	"fmt"

	"daydaynote/i18n"
	"daydaynote/storage"
	"daydaynote/utils"

	"github.com/spf13/cobra"
)

// delete 命令
func Delete() *cobra.Command {
	var dateValue string

	command := &cobra.Command{
		Use:   i18n.T(i18n.CmdDeleteUse),
		Short: i18n.T(i18n.CmdDeleteShort),
		Args:  utils.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			targetDate, err := utils.ResolveDate(dateValue)
			if err != nil {
				return err
			}

			index, err := utils.ParsePositiveIndex(args[0])
			if err != nil {
				return err
			}

			store := storage.StoreForDate(targetDate)
			records, err := store.List()
			if err != nil {
				return err
			}

			matchedPosition := 0
			deletePosition := -1
			for position, record := range records {
				if !utils.SameDate(record.CreatedAt.Local(), targetDate) {
					continue
				}

				matchedPosition++
				if matchedPosition == index {
					deletePosition = position
					break
				}
			}

			if deletePosition < 0 {
				return fmt.Errorf(i18n.T(i18n.ErrDeleteIndexNotFound), targetDate.Format("2006-01-02"), index)
			}

			updated := append(records[:deletePosition], records[deletePosition+1:]...)
			if err := store.Replace(updated); err != nil {
				return err
			}

			fmt.Printf(i18n.T(i18n.MsgDeleteSuccess), index, targetDate.Format("2006-01-02"))
			return nil
		},
	}

	command.Flags().StringVar(&dateValue, "date", "", i18n.T(i18n.FlagDate))
	return command
}
