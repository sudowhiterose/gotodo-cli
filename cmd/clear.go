package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sudowhiterose/gotodo-cli/internal/todo"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Удалить все выполненные задачи",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := todo.Load()
		if err != nil {
			return err
		}

		count := store.ClearDone()
		if err := store.Save(); err != nil {
			return err
		}

		fmt.Printf("Удалено выполненных задач: %d\n", count)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
