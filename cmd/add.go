/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	service "github.com/coderkhushal/proxabay/cmd/services"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds proxy on your device",
	Long: `run : proxabay add --origin url_of_your_server --port port_on_which_proxy_should_run_on_your_device
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n")
		pterm.DefaultBigText.WithLetters(putils.LettersFromString("Proxabay")).Render()

		origin, _ := cmd.Flags().GetString("origin")
		port, _ := cmd.Flags().GetString("port")

		if origin == "" {
			fmt.Println(service.Red, "URL is required", service.Reset)
			return
		}
		if port == "" {
			fmt.Println(service.Red, "port is required", service.Reset)
			return
		}
		pterm.NewRGB(0, 255, 255).Printfln("Port : %s , Origin : %s \n ", port, origin)

		err := ProxyManagerInstance.StartNewProxy(origin, ":"+port)
		if err != nil {
			pterm.Error.Println(err)
			return
		}
		pterm.Success.Println("Added proxy succesfully")

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
