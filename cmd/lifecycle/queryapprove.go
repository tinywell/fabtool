package lifecycle

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

func queryApprovedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "queryapprove",
		Short: "合约审批情况查询 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return queryApproved()
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
		"sequence",
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func queryApproved() error {
	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Channel:  channelID,
		Name:     chaincodeName,
		Sequence: chaincodeSequence,
		Peers:    peerAddresses,
	}
	info, err := service.LCQueryApproved(ccparam)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", info)
	return nil
}
