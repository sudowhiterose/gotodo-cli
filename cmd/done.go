package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sudowhiterose/gotodo-cli/internal/todo"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Отметить задачу выполненной",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("id должен быть числом")
		}

		store, err := todo.Load()
		if err != nil {
			return err
		}

		task, err := store.Find(id)
		if err != nil {
			return err
		}

		if task.Done {
			return fmt.Errorf("задача #%d уже выполнена", id)
		}

		task.Done = true
		if err := store.Save(); err != nil {
			return err
		}

		fmt.Printf("Задача #%d выполнена: %s\n", task.ID, task.Text)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
