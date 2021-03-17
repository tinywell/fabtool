package lifecycle

import (
	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

func commitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit",
		Short: "合约提交 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return commit()
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
		"channelID",
		"name",
		"version",
		"sequence",
		"isInit",
		"policy",
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func commit() error {
	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Channel:  channelID,
		Name:     chaincodeName,
		Version:  chaincodeVersion,
		Sequence: chaincodeSequence,
		Init:     isInit,
		Policy:   policy,
		Peers:    peerAddresses,
	}
	err := service.LCCommitCC(ccparam)
	if err != nil {
		return err
	}
	return nil
}
