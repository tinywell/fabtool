package lifecycle

import (
	"errors"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/tinywell/fabtool/pkg/sdk"
)

func packageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "package",
		Short: "合约打包 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 && len(args[0]) > 0 {
				return packageCC(args[0])
			}
			return errors.New("请指定打包文件名")
		},
	}
	flagList := []string{
		"lang",
		"path",
		"label",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func packageCC(fileName string) error {
	ccparam := sdk.CCParam{
		Path:  chaincodePath,
		Lang:  chaincodeLang,
		Label: chaincodeLabel,
	}
	pkg, err := sdk.LCPackageCC(ccparam)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, pkg, 0666)
	if err != nil {
		return err
	}
	return nil
}
