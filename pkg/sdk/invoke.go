package sdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// 合约执行方式
const (
	FCNInvoke = "invoke"
	FCNQuery  = "query"
)

// Invoke 合约执行
func (s *Service) Invoke(reqparam ExecuteReq) (*CCResult, error) {
	return s.execute(reqparam, FCNInvoke)
}

// Query 合约查询
func (s *Service) Query(reqparam ExecuteReq) (*CCResult, error) {
	return s.execute(reqparam, FCNQuery)
}

func (s *Service) execute(reqparam ExecuteReq, fcn string) (*CCResult, error) {
	ctx := s.sdk.ChannelContext(reqparam.Channel,
		fabsdk.WithUser(reqparam.User.User),
		fabsdk.WithOrg(reqparam.User.Org))
	client, err := channel.New(ctx)
	if err != nil {
		return nil, err
	}
	opts := []channel.RequestOption{}
	if len(reqparam.Peers) > 0 {
		opts = append(opts, channel.WithTargetEndpoints(reqparam.Peers...))
	}
	req := channel.Request{
		ChaincodeID:  reqparam.Chaincode,
		Fcn:          reqparam.Fcn,
		Args:         reqparam.Args,
		TransientMap: reqparam.TransientMap,
	}
	var resp channel.Response
	switch fcn {
	case FCNInvoke:
		resp, err = client.Execute(req, opts...)
	case FCNQuery:
		resp, err = client.Query(req, opts...)
	}
	if err != nil {
		return nil, err
	}
	return respToCCResult(resp), nil
}

func respToCCResult(resp channel.Response) *CCResult {
	return &CCResult{
		TxID:      string(resp.TransactionID),
		Status:    resp.ChaincodeStatus,
		ValidCode: resp.TxValidationCode.String(),
		Payload:   resp.Payload,
		Message:   resp.Responses[0].Response.Message,
	}
}
