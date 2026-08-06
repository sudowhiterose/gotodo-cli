package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: "Простой todo-менеджер в терминале",
	Long:  `gotodo — минималистичный cli для управления задачами. Данные хранятся в ~/.gotodo.json`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Здесь можно добавлять глобальные флаги
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}
