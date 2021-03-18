package chaincode

import (
	"encoding/json"
	"errors"
	"fmt"

	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/util"
	"github.com/spf13/cobra"

	"github.com/tinywell/fabtool/pkg/sdk"
)

func invokeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "合约调用",
		RunE: func(cmd *cobra.Command, args []string) error {
			return invoke()
		},
	}
	flagList := []string{
		"user",
		"org",
		"channelID",
		"name",
		"ctor",
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}
func queryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "合约查询",
		RunE: func(cmd *cobra.Command, args []string) error {
			return query()
		},
	}
	flagList := []string{
		"user",
		"org",
		"channelID",
		"name",
		"ctor",
		"peerAddresses",
	}
	attachFlags(cmd, flagList)
	return cmd
}

func invoke() error {
	return invokeOrQuery("invoke")
}
func query() error {
	return invokeOrQuery("query")
}

func invokeOrQuery(fcn string) error {
	err := serviceInit()
	if err != nil {
		return err
	}
	fmt.Println(chaincodeCtorJSON)
	input := &chaincodeInput{}
	if err := json.Unmarshal([]byte(chaincodeCtorJSON), &input); err != nil {
		return fmt.Errorf("解析 ctor 出错：%w", err)
	}
	var ccfcn string
	var args [][]byte
	if len(input.Args) > 0 {
		ccfcn = string(input.Args[0])
	} else {
		return errors.New("输入合约调用参数")
	}
	if len(input.Args) > 1 {
		args = input.Args[1:]
	}
	ccparam := sdk.ExecuteReq{
		User: sdk.Admin{
			User: user,
			Org:  org,
		},
		Channel:   channelID,
		Chaincode: chaincodeName,
		Fcn:       ccfcn,
		Args:      args,
		Peers:     peerAddresses,
	}
	var rest *sdk.CCResult
	switch fcn {
	case "invoke":
		rest, err = service.Invoke(ccparam)
	case "query":
		rest, err = service.Query(ccparam)
	}
	if err != nil {
		return err
	}
	resText := "合约执行结果：\n"
	resText += fmt.Sprintf("\tTxID: %s\n\tStatus: %d\n\tValid: %s\n\tMessage: %s\n\tPayload: %s",
		rest.TxID, rest.Status, rest.ValidCode, rest.Message, string(rest.Payload))
	fmt.Println(resText)
	return nil
}

type chaincodeInput struct {
	pb.ChaincodeInput
}

// UnmarshalJSON converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
func (c *chaincodeInput) UnmarshalJSON(b []byte) error {
	sa := struct {
		Function string
		Args     []string
	}{}
	err := json.Unmarshal(b, &sa)
	if err != nil {
		return err
	}
	allArgs := sa.Args
	if sa.Function != "" {
		allArgs = append([]string{sa.Function}, sa.Args...)
	}
	c.Args = util.ToChaincodeArgs(allArgs...)
	return nil
}
