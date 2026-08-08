/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/dakotaodev/brag/internal/brag"
	jsonstorage "github.com/dakotaodev/brag/internal/storage/json"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a brag!",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Incorrect number of arguments. ")
		}
		r, err := jsonstorage.NewJsonRepository()
		if err != nil {
			return err
		}
		entry := brag.Entry{
			ID:         uuid.NewString()[:6],
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
}
