/*
Copyright © 2023 NAME HERE <nik.datascience@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/nikitarudakov/tac/filenames"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")

		if err := filenames.RenameItemWithPath(args[0], args[1]); err != nil {
			log.Error().Str("path", args[0]).Err(err).Msg("error renaming item")
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
