package cmd

import (
	"errors"

	jsonstorage "github.com/dakotaodev/brag/internal/storage/json"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use: "update",
	Short: "Update an existing brag!",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Incorrect number of args.")
		}
		r, err := jsonstorage.NewJsonRepository()
		if err != nil {
			return err
		}		
		if err := r.Update(cmd.Context(), args[0], args[1]); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}