package chaincode

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/tinywell/fabtool/pkg/sdk"
)

func installCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "合约安装",
		RunE: func(cmd *cobra.Command, args []string) error {
			return install()
		},
	}
	flagList := []string{
		"user",
		"org",
		"lang",
		"path",
		"name",
		"version",
		"peerAddresses",
		"tlsRootCertFiles",
		"tarfile",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func install() error {
	err := serviceInit()
	if err != nil {
		return err
	}

	code, err := ioutil.ReadFile(chaincodeTarFile)
	if err != nil {
		return err
	}
	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Name:    chaincodeName,
		Version: chaincodeVersion,
		Path:    chaincodePath,
		Package: code,
		Lang:    chaincodeLang,
		Peers:   peerAddresses,
	}
	return service.InstallCC(ccparam)
}
