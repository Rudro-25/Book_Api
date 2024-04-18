/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Rudro-25/Book_API_Server/apiHandler"
	"github.com/spf13/cobra"
)

var Port int

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start cmd starts the server on a port",
	Long: `It starts the server on a given port number, 
				Port number will be given in the cmd`,
	Run: func(cmd *cobra.Command, args []string) {
		apiHandler.Start(Port)
	},
}

func init() {

	startCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8080, "default port for http server")
	rootCmd.AddCommand(startCmd)

}
