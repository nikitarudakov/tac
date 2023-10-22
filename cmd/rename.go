/*
Copyright © 2023 NAME HERE <nik.datascience@gmail.com>
*/
package cmd

import (
	"github.com/nikitarudakov/tac/filenames"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/spf13/cobra"
)

var (
	source string
	single bool
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files that follow some pattern",
	Long: `Rename command helps you manage naming of directories and files on 
your local computer. You can even specify patterns for naming`,
	Run: func(cmd *cobra.Command, args []string) {

		if single && source != "" {
			if err := filenames.RenameItemWithPath(source, args[0]); err != nil {
				log.Error().Str("path", source).Err(err).Msg("error renaming item")
			}

			return
		} else if source == "" {
			os.Exit(1)
		}

		if err := filenames.RenameItemsAtPath(source, "", []filenames.GroupRename{}); err != nil {
			log.Error().Str("path", source).Err(err).Msg("error renaming items")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	renameCmd.Flags().StringVar(
		&source,
		"src",
		"",
		"Set source of file(s) to rename")

	renameCmd.Flags().BoolVarP(
		&single,
		"single", "s",
		true,
		"Set to false to rename files in bulk")
}
