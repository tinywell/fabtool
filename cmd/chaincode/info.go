package chaincode

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

func infoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "合约部署情况查询",
		RunE: func(cmd *cobra.Command, args []string) error {
			return info()
		},
	}
	flagList := []string{
		"user",
		"org",
		"name",
		"peerAddresses",
		"tlsRootCertFiles",
		"channelID",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func info() error {
	err := serviceInit()
	if err != nil {
		return err
	}
	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Name:  chaincodeName,
		Peers: peerAddresses,
	}
	resp, err := service.QueryInstalledCC(ccparam)
	if err != nil {
		return err
	}
	fmt.Println(printResp(resp))
	return nil
}

func printResp(resp []*sdk.ChaincodeInfo) string {
	var respText string
	if len(resp) == 0 {
		respText = "节点上没有安装合约"
		return respText
	}
	respText = "节点上已安装合约有：\n"
	for i, r := range resp {
		rt := fmt.Sprintf("\t[%d]name:%s version:%s escc:%s path:%s\n", i+1, r.Name, r.Version, r.Escc, r.Path)
		respText += rt
	}
	return respText
}
