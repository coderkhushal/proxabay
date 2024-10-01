/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	server "github.com/coderkhushal/proxabay/cmd/proxyserver"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "proxabay",
	Short: "Proxabay is a reverse proxy with caching",
	Long: `Proxabay is a reverse proxy with caching that can be used to cache the responses from the upstream server.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	go func() {

		err := rootCmd.Execute()
		if err != nil {
			os.Exit(1)
		}
	}()

	waitforShutDown()
}

func waitforShutDown() {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("press ctr+c to continue")
	<-sigch
	fmt.Println("process ended!!")
}

var ProxyManagerInstance *server.ProxyManger = server.NewProxyManger()

func init() {

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.proxabay.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
