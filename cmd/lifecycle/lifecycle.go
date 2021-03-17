package lifecycle

import (
	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

var (

	// LifecycleCMD ...
	LifecycleCMD = cobra.Command{
		Use:   "lifecycle",
		Short: "合约管理 - lifecycle",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	service *sdk.Service
	config  string
)

// Cmd ...
func Cmd() *cobra.Command {
	LifecycleCMD.AddCommand(packageCmd())
	LifecycleCMD.AddCommand(installCmd())
	LifecycleCMD.AddCommand(queryInstallCmd())
	LifecycleCMD.AddCommand(getInstallCmd())
	LifecycleCMD.AddCommand(approveCmd())
	LifecycleCMD.AddCommand(queryApprovedCmd())
	LifecycleCMD.AddCommand(checkCommitCmd())
	LifecycleCMD.AddCommand(commitCmd())
	LifecycleCMD.AddCommand(queryCommitedCmd())
	LifecycleCMD.PersistentFlags().StringVarP(&config, "config", "c", "config.yaml", "配置文件路径")
	return &LifecycleCMD
}
