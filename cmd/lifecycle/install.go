package lifecycle

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/tinywell/fabtool/pkg/sdk"
)

func installCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "合约安装 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 && len(args[0]) > 0 {
				return install(args[0])
			}
			return errors.New("请指定安装的合约打包文件")
		},
	}
	flagList := []string{
		"user",
		"org",
		"label",
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func install(fileName string) error {
	err := serviceInit()
	if err != nil {
		return err
	}
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Label:   chaincodeLabel,
		Package: code,
		Peers:   peerAddresses,
	}
	packageID, err := service.LCInstallCC(ccparam)
	if err != nil {
		return err
	}
	fmt.Printf("合约安装完成，package ID = %s\n", packageID)
	return nil
}
