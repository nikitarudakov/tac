/*
Copyright © 2023 NAME HERE <nik.datascience@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var defaultLogger zerolog.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tac",
	Short: "CLI for managing files on your local PC",
	Long: `Welcome to Tac - your go-to tool for managing files easily via the command line. 

Use short commands for file decompression, renaming, and format conversion. 

Simplify complex tasks and boost your productivity effortlessly.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	zerolog.SetGlobalLevel(-1)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Trace().
		Int("logger_level", -1).
		Msg("logger was set")

	defaultLogger.Info().Msg("Logger is running ...")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config path (default is $HOME/.tac.json)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config zip file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".Traffic-Analyst-CLI" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".tac")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config zipfile is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config zipfile:", viper.ConfigFileUsed())
	}
}
