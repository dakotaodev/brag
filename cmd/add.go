/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/dakotaodev/brag/internal/brag"
	jsonstorage "github.com/dakotaodev/brag/internal/storage/json"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a win to your brag document.",
	Long:  `A longer description`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			fmt.Printf("Incorrect number of arguments. You provided %d, only 1 is expected", len(args))
		}
		r := jsonstorage.NewJsonRepository("./brag.json")
		entry := brag.Entry{
			Value:      args[0],
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}
		if _, err := r.Add(cmd.Context(), entry); err != nil {
			return fmt.Errorf("add brag entry: %w", err)
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
