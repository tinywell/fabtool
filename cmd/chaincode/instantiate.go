package chaincode

import (
	"github.com/spf13/cobra"
	"github.com/tinywell/fabtool/pkg/sdk"
)

func instantiateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "instantiate",
		Short: "合约实例化",
		RunE: func(cmd *cobra.Command, args []string) error {
			return instantiate()
		},
	}
	flagList := []string{
		"user",
		"org",
		"channelID",
		"lang",
		"path",
		"name",
		"version",
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func instantiate() error {
	err := serviceInit()
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
		Lang:    chaincodeLang,
		Channel: channelID,
	}
	return service.InstantiateCC(ccparam)
}
