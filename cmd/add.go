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

// add 命令
func Add() *cobra.Command {
	return &cobra.Command{
		Use:   i18n.T(i18n.CmdAddUse),
		Short: i18n.T(i18n.CmdAddShort),
		Args:  utils.MinimumNArgs(1),
		RunE:  AddHandle,
	}
}

// 把命令行传入的内容拼成一条记录并追加写入存储
func AddHandle(cmd *cobra.Command, args []string) error {
	content := strings.TrimSpace(strings.Join(args, " "))
	if content == "" {
		return errors.New(i18n.T(i18n.ErrAddEmpty))
	}

	now := time.Now()
	record := storage.Record{
		ID:        fmt.Sprintf("%d", now.UnixNano()),
		Title:     content,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}

	store := storage.DefaultStore()
	if err := store.Append(record); err != nil {
		return err
	}

	fmt.Printf(i18n.T(i18n.MsgAddSaved), store.Path())
	return nil
}
