/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"syscall"

	service "github.com/coderkhushal/proxabay/cmd/services"
	"github.com/spf13/cobra"
)

// clearcacheCmd represents the clearcache command
var clearcacheCmd = &cobra.Command{
	Use:   "clearcache",
	Short: "Clears the cache of requests",
	Long:  `Run: proxabay clearcache`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ClearCache()
		service.Sigch <- syscall.SIGINT

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
