package lifecycle

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

func queryInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "queryinstalled",
		Short: "已安装合约查询 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return queryInstall()
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
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func queryInstall() error {
	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Peers: peerAddresses,
	}
	ccinfo, err := service.LCQueryInstalledCC(ccparam)
	if err != nil {
		return err
	}
	fmt.Println(printResp(ccinfo))
	return nil
}

func printResp(resp []*sdk.LCChaincodeInfo) string {
	var respText string
	if len(resp) == 0 {
		respText = "节点上没有安装合约"
		return respText
	}
	respText = "节点上已安装合约有：\n"
	for i, r := range resp {
		rt := fmt.Sprintf("\t[%d]Package ID:%s Label:%s References:%+v \n", i+1, r.PackageID, r.Label, r.References)
		respText += rt
	}
	return respText
}
