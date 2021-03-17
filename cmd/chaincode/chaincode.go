package chaincode

import (
	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

var (

	// ChaincodeCMD ...
	ChaincodeCMD = cobra.Command{
		Use:   "chaincode",
		Short: "合约管理",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	service *sdk.Service
	config  string
)

// Cmd ...
func Cmd() *cobra.Command {
	ChaincodeCMD.AddCommand(installCmd())
	ChaincodeCMD.AddCommand(infoCmd())
	ChaincodeCMD.AddCommand(instantiateCmd())
	ChaincodeCMD.PersistentFlags().StringVarP(&config, "config", "c", "config.yaml", "配置文件路径")
	return &ChaincodeCMD
}
