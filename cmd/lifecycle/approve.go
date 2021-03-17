package lifecycle

import (
	"github.com/spf13/cobra"
	"github.com/tinywell/fabtool/pkg/sdk"
)

func approveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve",
		Short: "合约审批 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return approve()
		},
	}
	flagList := []string{
		"user",
		"org",
		"channelID",
		"name",
		"version",
		"sequence",
		"packageID",
		"isInit",
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func approve() error {
	err := serviceInit()
	if err != nil {
		return err
	}

	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Channel:   channelID,
		Name:      chaincodeName,
		Version:   chaincodeVersion,
		PackageID: chaincodePKGID,
		Sequence:  chaincodeSequence,
		Policy:    policy,
		Init:      isInit,
		Peers:     peerAddresses,
	}
	err = service.LCApproveCC(ccparam)
	if err != nil {
		return err
	}
	return nil
}
