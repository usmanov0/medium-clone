package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "grpc-server",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("hello world")

		c := make(chan string)
		<-c
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
