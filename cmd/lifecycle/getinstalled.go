package lifecycle

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/tinywell/fabtool/pkg/sdk"
)

func getInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getinstalled",
		Short: "已安装合约包下载 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getInstalled()
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := serviceInit()
			if err != nil {
				return err
			}
			return nil
		},
	}
	flagList := []string{
		"user",
		"org",
		"packageID",
		"peerAddresses",
		"output",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func getInstalled() error {
	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Peers:     peerAddresses,
		PackageID: chaincodePKGID,
	}
	code, err := service.LCGetInstalledCC(ccparam)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(output, code, 0666)
}
