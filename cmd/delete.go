package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sudowhiterose/gotodo-cli/internal/todo"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Удалить задачу",
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

		if err := store.Delete(id); err != nil {
			return err
		}

		if err := store.Save(); err != nil {
			return err
		}

		fmt.Printf("Задача #%d удалена\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
