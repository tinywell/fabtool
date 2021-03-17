package lifecycle

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

func queryCommitedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "querycommit",
		Short: "合约提交情况查询 - lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			return queryCommitted()
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
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func queryCommitted() error {
	ccparam := sdk.CCParam{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Channel: channelID,
		Name:    chaincodeName,
		Peers:   peerAddresses,
	}
	resp, err := service.LCQueryCommitted(ccparam)
	if err != nil {
		return err
	}
	for i, ci := range resp {
		fmt.Printf("[%d] %+v\n", i, ci)
	}
	return nil
}
