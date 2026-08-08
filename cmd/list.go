/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	jsonstorage "github.com/dakotaodev/brag/internal/storage/json"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List exiting brags!",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := jsonstorage.NewJsonRepository()
		if err != nil {
			return err
		}
		entries, err := r.List(cmd.Context())
		if err != nil {
			return err
		}
		if len(entries) == 0 {
			fmt.Println("No existing entries, find something to brag about!")
			return nil
		}
		fmt.Println("===========================")
		for i := 0; i < len(entries); i++ {
			e := entries[i]
			fmt.Printf("%s - %s - %s\n", e.ID, e.CreatedAt.Format("Jan 2, 2006 at 3:04 PM"), e.Value)
		}
		fmt.Println("===========================")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
