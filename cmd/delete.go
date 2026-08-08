package cmd

import (
	"errors"

	jsonstorage "github.com/dakotaodev/brag/internal/storage/json"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete a brag!",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Incorrect number of args.")
		}
		r:=jsonstorage.NewJsonRepository("./brag.json")
		if err:= r.Delete(cmd.Context(), args[0]); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}