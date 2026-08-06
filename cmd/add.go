package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sudowhiterose/gotodo-cli/internal/todo"
)

var addCmd = &cobra.Command{
	Use:   "add [текст задачи]",
	Short: "Добавить новую задачу",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := strings.Join(args, " ")
		if strings.TrimSpace(text) == "" {
			return fmt.Errorf("текст задачи не может быть пустым")
		}

		store, err := todo.Load()
		if err != nil {
			return err
		}

		task := store.Add(text)
		if err := store.Save(); err != nil {
			return err
		}

		fmt.Printf("Добавлена задача #%d: %s\n", task.ID, task.Text)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
