/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		origin, _ := cmd.Flags().GetString("origin")
		port, _ := cmd.Flags().GetString("port")

		if origin == "" {
			log.Fatal("URL is required")
			return
		}
		if port == "" {
			log.Fatal("port is required")
			return
		}
		fmt.Printf("port : %s , origin : %s \n", port, origin)
		err := ProxyManagerInstance.StartNewProxy(origin, ":"+port)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Added proxy succesfully")

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().String("port", "", "Port of server")
	addCmd.PersistentFlags().String("origin", "", "Origin URL")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
