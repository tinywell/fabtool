package lifecycle

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

func checkCommitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checkcommit",
		Short: "合约提交前置检查 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkCommit()
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

func checkCommit() error {
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
	resp, err := service.LCCheckCommit(ccparam)
	if err != nil {
		return err
	}
	for _, ap := range resp {
		fmt.Println(ap)
	}
	return nil
}
