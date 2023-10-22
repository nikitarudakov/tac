/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nikitarudakov/tac/zipfile"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	zip    bool
	create bool
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "decompress",
	Short: "Decompress files from various formats",
	Long: `The decompress command in Tac allows you to effortlessly extract compressed files 
from a diverse range of formats, including but not limited to zip, tar, gzip, and more.`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Unzip command was called")

		if len(args) > 1 {
			in := zipfile.Input{}
			in.SetInput(args[0], zip)

			out := zipfile.Output{}
			out.SetOutput(args[1], create)

			if err := zipfile.Decompress(in, out); err != nil {
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	extractCmd.Flags().BoolVarP(
		&zip,
		"zip", "z",
		false, "use to decompress zip files (default is false)",
	)

	extractCmd.Flags().BoolVarP(
		&zip,
		"create", "c",
		false, "use to create output folder (default is false)",
	)
}
