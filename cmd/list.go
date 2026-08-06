package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/sudowhiterose/gotodo-cli/internal/todo"
)

var (
	showAll bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать список задач",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := todo.Load()
		if err != nil {
			return err
		}

		if len(store.Tasks) == 0 {
			fmt.Println("Задач пока нет. Добавь первую: gotodo add \"...\"")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tСТАТУС\tСОЗДАНА\tЗАДАЧА")

		for _, t := range store.Tasks {
			if !showAll && t.Done {
				continue
			}
			status := "[ ]"
			if t.Done {
				status = "[✓]"
			}
			created := t.CreatedAt.Format("02.01 15:04")
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", t.ID, status, created, t.Text)
		}
		return w.Flush()
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "показывать и выполненные задачи")
	rootCmd.AddCommand(listCmd)
}
