/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	service "github.com/coderkhushal/proxabay/cmd/services"
	"github.com/spf13/cobra"
)

// clearcacheCmd represents the clearcache command
var clearcacheCmd = &cobra.Command{
	Use:   "clearcache",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ClearCache()
		fmt.Println("Cache cleared")

	},
}

func init() {
	rootCmd.AddCommand(clearcacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearcacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearcacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
