package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/cmd/chaincode"
	"github.com/tinywell/fabtool/cmd/lifecycle"
)

var (
	// RootCMD 根
	RootCMD = cobra.Command{
		Use:   "app",
		Short: "fabirc 客户端工具",
		Run: func(*cobra.Command, []string) {

		},
	}
	config string
)

func init() {

}

func main() {

	RootCMD.AddCommand(chaincode.Cmd())
	RootCMD.AddCommand(lifecycle.Cmd())
	if err := RootCMD.Execute(); err != nil {
		os.Exit(1)
	}
}
